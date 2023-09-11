package routers

import (
	"net/http"
	"time"

	_ "github.com/haierspi/golang-image-upload-service/docs"
	"github.com/haierspi/golang-image-upload-service/global"
	"github.com/haierspi/golang-image-upload-service/internal/middleware"
	"github.com/haierspi/golang-image-upload-service/internal/routers/api"

	"github.com/haierspi/golang-image-upload-service/pkg/limiter"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.BucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.AppInfo())

	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	//对404 的处理
	r.NoRoute(middleware.NoFound())
	//	r.Use(middleware.Tracing())
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())
	r.Use(middleware.Cors())
	//r.Use(middleware.AuthToken())

	r.GET("/debug/vars", api.Expvar)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	upload := api.NewUpload()
	r.Group("/api").Use(middleware.AuthToken()).POST("/upload", upload.UploadFile)

	//r.Use(middleware.AuthToken()).POST("/upload", upload.UploadFile)

	r.StaticFS(global.AppSetting.UploadSavePath, http.Dir(global.AppSetting.UploadSavePath))

	return r
}
