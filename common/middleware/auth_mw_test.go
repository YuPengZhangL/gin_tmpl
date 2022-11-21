package middleware

import (
	"gg_web_tmpl/common/config"
	"gg_web_tmpl/common/log"
	"gg_web_tmpl/common/utils"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	log.InitMockLogger()
	config.InitMockConfig()
}

func TestAuthMW(t *testing.T) {
	Convey("Test AuthMW", t, func() {
		Convey("should success", func() {
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.Use(AuthMW())
			r.GET("/test", func(c *gin.Context) {
				id, _ := c.Get("user_id")
				idStr := id.(string)
				c.String(http.StatusOK, idStr)
			})
			req, _ := http.NewRequest("GET", "/test", nil)
			token, _ := utils.GenerateToken(1)
			req.Header.Set("Authorization", token)
			r.ServeHTTP(w, req)
			So(w.Body.String(), ShouldEqual, "1")
		})
		Convey("should fail: token is empty", func() {
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.Use(AuthMW())
			r.GET("/test", func(c *gin.Context) {
				id, _ := c.Get("user_id")
				idStr := id.(string)
				c.String(http.StatusOK, idStr)
			})
			req, _ := http.NewRequest("GET", "/test", nil)
			r.ServeHTTP(w, req)
			So(w.Body.String(), ShouldEqual, `{"status_code":-3,"status_msg":"that's not even a token"}`)
		})
		Convey("should fail: token error", func() {
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.Use(AuthMW())
			r.GET("/test", func(c *gin.Context) {
				id, _ := c.Get("user_id")
				idStr := id.(string)
				c.String(http.StatusOK, idStr)
			})
			req, _ := http.NewRequest("GET", "/test", nil)
			req.Header.Set("Authorization", "test")
			r.ServeHTTP(w, req)
			So(w.Body.String(), ShouldEqual, `{"status_code":-3,"status_msg":"that's not even a token"}`)
		})
	})
}
