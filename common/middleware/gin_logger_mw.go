package middleware

import (
	"gg_web_tmpl/common/log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httputil"
	"runtime/debug"
	"time"
)

// GinLoggerMW gin日志记录中间件
func GinLoggerMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录开始时间
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		logger := log.GetLoggerWithCtx(c)
		// 结束时间
		end := time.Now()
		// 接口耗时
		latency := end.Sub(start)
		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				logger.Error(e)
			}
		} else {
			logger.WithFields(logrus.Fields{
				"status":     c.Writer.Status(),
				"method":     c.Request.Method,
				"path":       path,
				"query":      query,
				"ip":         c.ClientIP(),
				"user-agent": c.Request.UserAgent(),
				"latency":    latency,
			}).Infof("success")
		}
	}
}

// GinRecoverLoggerMW gin panic recover中间件
func GinRecoverLoggerMW(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger := log.GetLoggerWithCtx(c)
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				path := c.Request.URL.Path
				method := c.Request.Method
				if stack {
					logger.WithFields(logrus.Fields{
						"method":  method,
						"path":    path,
						"request": string(httpRequest),
						"error":   err,
						"stack":   string(debug.Stack()),
					}).Error("[Recovery from panic]")
				} else {
					logger.WithFields(logrus.Fields{
						"method":  method,
						"path":    path,
						"request": string(httpRequest),
						"error":   err,
					}).Error("[Recovery from panic]")
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
