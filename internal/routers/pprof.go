package routers

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/haierspi/golang-image-upload-service/global"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	// DefaultPrefix url prefix of pprof
	DefaultPrefix = "/debug/pprof"
)

func MetricsSrv() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("PrivateHttp Metric Service panic", "err", err)
		}
	}()

	log.Println("PrivateHttp Metric Service ListenAndServe On: ", global.ServerSetting.PrivateHttpListen, "\n")

	router := gin.New()
	//router.Use(logs.RecoveryWithZap(logs.Logger, true))

	//prom监控
	router.GET("metrics", gin.WrapH(promhttp.Handler()))

	if global.ServerSetting.RunMode == "debug" {
		registerPprof(router)
	}

	s := &http.Server{
		Addr:           ":" + global.ServerSetting.PrivateHttpListen,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	fmt.Println("PrivateHttp Metric Service start err", "err", err)
}

func registerPprof(r *gin.Engine, prefixOptions ...string) {
	prefix := getPrefix(prefixOptions...)

	prefixRouter := r.Group(prefix)
	{
		prefixRouter.GET("/", pprofHandler(pprof.Index))
		prefixRouter.GET("/cmdline", pprofHandler(pprof.Cmdline))
		prefixRouter.GET("/profile", pprofHandler(pprof.Profile))
		prefixRouter.POST("/symbol", pprofHandler(pprof.Symbol))
		prefixRouter.GET("/symbol", pprofHandler(pprof.Symbol))
		prefixRouter.GET("/trace", pprofHandler(pprof.Trace))
		prefixRouter.GET("/allocs", pprofHandler(pprof.Handler("allocs").ServeHTTP))
		prefixRouter.GET("/block", pprofHandler(pprof.Handler("block").ServeHTTP))
		prefixRouter.GET("/goroutine", pprofHandler(pprof.Handler("goroutine").ServeHTTP))
		prefixRouter.GET("/heap", pprofHandler(pprof.Handler("heap").ServeHTTP))
		prefixRouter.GET("/mutex", pprofHandler(pprof.Handler("mutex").ServeHTTP))
		prefixRouter.GET("/threadcreate", pprofHandler(pprof.Handler("threadcreate").ServeHTTP))
	}
}
func pprofHandler(h http.HandlerFunc) gin.HandlerFunc {
	handler := h
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
func getPrefix(prefixOptions ...string) string {
	prefix := DefaultPrefix
	if len(prefixOptions) > 0 {
		prefix = prefixOptions[0]
	}
	return prefix
}
