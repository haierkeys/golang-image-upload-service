package main

import (
	"context"
	_ "expvar"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	validatorV10 "github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/haierspi/golang-image-upload-service/global"
	"github.com/haierspi/golang-image-upload-service/internal/routers"
	"github.com/haierspi/golang-image-upload-service/pkg/logger"
	"github.com/haierspi/golang-image-upload-service/pkg/setting"
	"github.com/haierspi/golang-image-upload-service/pkg/tracer"
	"github.com/haierspi/golang-image-upload-service/pkg/validator"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	port      string
	runMode   string
	config    string
	Version   string
	isVersion bool
	GitTag    = "2000.01.01.release"
	BuildTime = "2000-01-01T00:00:00+0800"
)

func init() {
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}

	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	err = setupValidator()
	if err != nil {
		log.Fatalf("init.setupValidator err: %v", err)
	}

	//err = setupTracer()
	//if err != nil {
	//	log.Fatalf("init.setupTracer err: %v", err)
	//}
}

func main() {

	if isVersion {
		fmt.Println("Version: " + Version)
		fmt.Println("Git Tag: " + GitTag)
		fmt.Println("Build Time: " + BuildTime)
		return
	}

	fmt.Println("Config Path : " + global.CONFIG)

	gin.SetMode(global.ServerSetting.RunMode)

	validator.RegisterCustom()

	go routers.MetricsSrv()

	log.Println("Http API Service ListenAndServe On: ", global.ServerSetting.HttpPort, "\n")
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {

		if err := s.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("s.ListenAndServe err: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shuting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/config.yaml", "指定要使用的配置文件路径")
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	v := flag.Bool("v", false, "编译信息缩写")
	flag.Parse()
	if *v {
		isVersion = true
	}
	return nil
}

func setupSetting() error {
	global.CONFIG = config

	s, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("OSS", &global.OSSSetting)
	if err != nil {
		return err
	}

	global.AppSetting.DefaultContextTimeout *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	mw := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	})

	global.Logger = logger.NewLogger(mw, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupValidator() error {
	global.Validator = validator.NewCustomValidator()
	global.Validator.Engine()
	binding.Validator = global.Validator
	uni := ut.New(en.New(), en.New(), zh.New())
	v, ok := binding.Validator.Engine().(*validatorV10.Validate)
	if ok {
		zhTran, _ := uni.GetTranslator("zh")
		enTran, _ := uni.GetTranslator("en")
		err := zhTranslations.RegisterDefaultTranslations(v, zhTran)
		if err != nil {
			return err
		}
		err = enTranslations.RegisterDefaultTranslations(v, enTran)
		if err != nil {
			return err
		}
	}

	global.Ut = uni

	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blog-service", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
