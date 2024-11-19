package middleware

import (
	"github.com/haierspi/golang-image-upload-service/global"
	"github.com/haierspi/golang-image-upload-service/pkg/app"

	"github.com/gin-gonic/gin"
)

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", global.Name)
		c.Set("app_version", global.Version)
		c.Set("access_host", app.GetAccessHost(c))

		c.Next()
	}
}
