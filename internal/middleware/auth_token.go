/**
  @author: HaierSpi
  @since: 2022/9/14
  @desc:
**/

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/haierspi/golang-image-upload-service/global"
	"github.com/haierspi/golang-image-upload-service/pkg/app"
	"github.com/haierspi/golang-image-upload-service/pkg/code"
)

func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		if global.SecuritySetting.AuthToken == "" {
			c.Next()
		}

		response := app.NewResponse(c)

		var token string

		if s, exist := c.GetQuery("authorization"); exist {
			token = s
		} else if s, exist = c.GetQuery("Authorization"); exist {
			token = s
		} else if s = c.GetHeader("authorization"); len(s) != 0 {
			token = s
		} else if s = c.GetHeader("Authorization"); len(s) != 0 {
			token = s
		}

		if token != global.SecuritySetting.AuthToken {
			response.ToResponse(code.ErrorInvalidAuthToken)
			c.Abort()
		}
		c.Next()
	}
}
