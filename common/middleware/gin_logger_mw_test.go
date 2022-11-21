package middleware

import (
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGinLoggerMW(t *testing.T) {
	Convey("Test GinLoggerMW", t, func() {
		w := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		_, r := gin.CreateTestContext(w)
		r.Use(GinLoggerMW())
		r.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "ok")
		})
		req, _ := http.NewRequest("GET", "/test", nil)
		r.ServeHTTP(w, req)
		So(w.Body.String(), ShouldEqual, "ok")
	})
}

func TestGinRecoverLoggerMW(t *testing.T) {
	Convey("Test GinRecoverLoggerMW", t, func() {
		Convey("handler panic: print stack", func() {
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.Use(GinRecoverLoggerMW(true))
			r.GET("/test", func(c *gin.Context) {
				panic("test panic")
			})
			req, _ := http.NewRequest("GET", "/test", nil)
			r.ServeHTTP(w, req)
			So(w.Code, ShouldEqual, http.StatusInternalServerError)
		})
		Convey("handler panic: no print stack", func() {
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.Use(GinRecoverLoggerMW(false))
			r.GET("/test", func(c *gin.Context) {
				panic("test panic")
			})
			req, _ := http.NewRequest("GET", "/test", nil)
			r.ServeHTTP(w, req)
			So(w.Code, ShouldEqual, http.StatusInternalServerError)
		})
		Convey("handler normal", func() {
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.Use(GinRecoverLoggerMW(false))
			r.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "ok")
			})
			req, _ := http.NewRequest("GET", "/test", nil)
			r.ServeHTTP(w, req)
			So(w.Body.String(), ShouldEqual, "ok")
			So(w.Code, ShouldEqual, http.StatusOK)
		})
	})
}
