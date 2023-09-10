package middleware

import (
	"github.com/haierspi/golang-image-upload-service/pkg/app"
	"github.com/haierspi/golang-image-upload-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)

		if s, exist := c.GetQuery("authorization"); exist {
			token = s
		} else if s, exist = c.GetQuery("Authorization"); exist {
			token = s
		} else if s, exist := c.GetQuery("token"); exist {
			token = s
		} else if s, exist = c.GetQuery("Token"); exist {
			token = s
		} else if s = c.GetHeader("authorization"); len(s) != 0 {
			token = s
		} else if s = c.GetHeader("Authorization"); len(s) != 0 {
			token = s
		} else if s = c.GetHeader("token"); len(s) != 0 {
			token = s
		} else if s = c.GetHeader("Token"); len(s) != 0 {
			token = s
		}

		if token == "" {
			ecode = errcode.InvalidToken
		} else {
			if user, err := app.ParseToken(token); err != nil {
				ecode = errcode.UnauthorizedTokenError
			} else {
				c.Set("user_token", user)
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return

		}

		c.Next()
	}
}
