package routers

import (
	"time"

	_ "github.com/haierspi/golang-image-upload-service/docs"
	"github.com/haierspi/golang-image-upload-service/global"
	"github.com/haierspi/golang-image-upload-service/internal/middleware"
	"github.com/haierspi/golang-image-upload-service/internal/routers/api"
	"github.com/haierspi/golang-image-upload-service/internal/routers/api/v1"
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

	r.GET("/debug/vars", api.Expvar)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//upload := api.NewUpload()
	//r.POST("/upload/file", upload.UploadFile)

	//r.StaticFS("/static", httpclient.Dir(global.AppSetting.UploadSavePath))

	// //middleware.JWT() middleware.Token()
	//APIV1Token := r.Group("/api/v1").Use(middleware.Token())

	APIV1 := r.Group("/api/v1")

	goods := v1.NewGoods()
	APIV1.Use()
	{
		//分类
		APIV1.GET("/category/list", goods.List)
		APIV1.GET("/category/details", goods.Details)

	}

	client := v1.NewClient()
	APIV1.GET("/client/version", client.Version)

	return r
}
