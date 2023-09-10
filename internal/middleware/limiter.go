package middleware

import (
	"github.com/haierspi/golang-image-upload-service/pkg/app"
	"github.com/haierspi/golang-image-upload-service/pkg/errcode"
	"github.com/haierspi/golang-image-upload-service/pkg/limiter"

	"github.com/gin-gonic/gin"
)

func RateLimiter(l limiter.Face) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
