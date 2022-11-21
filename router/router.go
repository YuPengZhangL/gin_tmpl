package router

import (
	"gg_web_tmpl/common/config"
	"gg_web_tmpl/common/middleware"
	"gg_web_tmpl/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	if config.GetConf().App.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(middleware.RequestIdMW())
	r.Use(middleware.CorsMW())
	r.Use(middleware.GinLoggerMW())
	r.Use(middleware.GinRecoverLoggerMW(true))
	r.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "gg_web_tmpl API is running...")
	})
	setupUserRouter(r)
	setupUserAuthRouter(r)
	return r
}

func setupUserRouter(r *gin.Engine) {
	r.POST("/api/v1/user/register", handler.RegisterUser)
	r.POST("/api/v1/user/login", handler.Login)
}

func setupUserAuthRouter(r *gin.Engine) {
	userGroup := r.Group("/api/v1/user")
	userGroup.Use(middleware.AuthMW())
	userGroup.GET("/info", handler.GetUserInfo)
}
