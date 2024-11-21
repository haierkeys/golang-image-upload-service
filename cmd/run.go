package cmd

import (
    "context"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "os/signal"
    "reflect"
    "strings"
    "syscall"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/locales/en"
    "github.com/go-playground/locales/zh"
    ut "github.com/go-playground/universal-translator"
    validatorV10 "github.com/go-playground/validator/v10"
    en_translations "github.com/go-playground/validator/v10/translations/en"
    zh_translations "github.com/go-playground/validator/v10/translations/zh"
    "github.com/pkg/errors"
    "github.com/spf13/cobra"
    "go.uber.org/zap"
    "gopkg.in/natefinch/lumberjack.v2"

    "github.com/haierspi/golang-image-upload-service/global"
    "github.com/haierspi/golang-image-upload-service/internal/routers"
    "github.com/haierspi/golang-image-upload-service/pkg/logger"
    "github.com/haierspi/golang-image-upload-service/pkg/path"
    "github.com/haierspi/golang-image-upload-service/pkg/validator"
)

type runFlags struct {
    dir     string // 项目根目录
    port    string // 启动端口
    runMode string // 启动模式
    config  string // 指定要使用的配置文件路径
}

func init() {
    runEnv := new(runFlags)

    var runCommand = &cobra.Command{
        Use:   "run [-c config_file] [-d working_dir] [-p port]",
        Short: "Run service",
        Run: func(cmd *cobra.Command, args []string) {

            if len(runEnv.dir) > 0 {
                err := os.Chdir(runEnv.dir)
                if err != nil {
                    fmt.Println("failed to change the current working directory, ", err)
                }
                fmt.Println("working directory changed", zap.String("path", runEnv.dir).String)
            }

            if len(runEnv.config) <= 0 {
                if path.Exists("config.yaml") {
                    runEnv.config = "config.yaml"
                } else if path.Exists("config/config-dev.yaml") {
                    runEnv.config = "config/config-dev.yaml"
                } else {
                    runEnv.config = "config/config.yaml"
                }
            }

            if err := global.ConfigLoad(runEnv.config); err != nil {
                fmt.Println(err)
            }

            if len(runEnv.runMode) > 0 {
                global.Config.Server.RunMode = runEnv.runMode
            }

            fmt.Println("Config Path : " + runEnv.config)

            gin.SetMode(global.Config.Server.RunMode)

            initLogger()
            initValidator()

            validator.RegisterCustom()

            go routers.MetricsSrv()

            log.Println("Http API Service ListenAndServe On: ", global.Config.Server.HttpPort, "\n")
            router := routers.NewRouter()
            s := &http.Server{
                Addr:           global.Config.Server.HttpPort,
                Handler:        router,
                ReadTimeout:    time.Duration(global.Config.Server.ReadTimeout) * time.Second,
                WriteTimeout:   time.Duration(global.Config.Server.WriteTimeout) * time.Second,
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

        },
    }

    rootCmd.AddCommand(runCommand)
    fs := runCommand.Flags()
    fs.StringVarP(&runEnv.dir, "dir", "d", "", "run dir")
    fs.StringVarP(&runEnv.port, "port", "p", "", "run port")
    fs.StringVarP(&runEnv.runMode, "mode", "m", "", "run mode")
    fs.StringVarP(&runEnv.config, "config", "c", "", "config file")

}

func initLogger() error {
    fileName := global.Config.App.LogSavePath + "/" + global.Config.App.LogFile

    mw := io.MultiWriter(os.Stdout, &lumberjack.Logger{
        Filename:  fileName,
        MaxSize:   500,
        MaxAge:    10,
        LocalTime: true,
    })

    global.Logger = logger.NewLogger(mw, "", log.LstdFlags)

    return nil
}

func initValidator() error {
    global.Validator = validator.NewCustomValidator()
    global.Validator.Engine()
    binding.Validator = global.Validator

    var uni *ut.UniversalTranslator

    validate, ok := binding.Validator.Engine().(*validatorV10.Validate)
    if ok {

        validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
            name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
            if name == "-" {
                return ""
            }
            return name
        })

        uni = ut.New(en.New(), en.New(), zh.New())

        zhTran, _ := uni.GetTranslator("zh")
        enTran, _ := uni.GetTranslator("en")

        err := zh_translations.RegisterDefaultTranslations(validate, zhTran)
        if err != nil {
            return err
        }
        err = en_translations.RegisterDefaultTranslations(validate, enTran)
        if err != nil {
            return err
        }
    }

    global.Ut = uni

    return nil
}
