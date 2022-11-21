package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// RequestIdMW 生成RequestID中间件
func RequestIdMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get("X-Request-Id")
		if requestId == "" {
			reqUUID, _ := uuid.NewV4()
			requestId = reqUUID.String()
			c.Request.Header.Set("X-Request-Id", requestId)
			c.Set("request_id", requestId)
		}
		c.Header("X-Request-Id", requestId)
		c.Next()
	}
}
