package middleware

import (
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCorsMW(t *testing.T) {
	Convey("Test CrosMW", t, func() {
		Convey("normal", func() {
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.Use(CorsMW())
			r.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "ok")
			})
			req, _ := http.NewRequest("GET", "/test", nil)
			r.ServeHTTP(w, req)
			So(w.Body.String(), ShouldEqual, "ok")
			So(w.Header().Get("Access-Control-Allow-Origin"), ShouldEqual, "*")
			So(w.Header().Get("Access-Control-Allow-Headers"), ShouldEqual, "*")
			So(w.Header().Get("Access-Control-Allow-Methods"), ShouldEqual, "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			So(w.Header().Get("Access-Control-Expose-Headers"), ShouldEqual, "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			So(w.Header().Get("Access-Control-Allow-Credentials"), ShouldEqual, "true")
		})
		Convey("options method", func() {
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.Use(CorsMW())
			r.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "ok")
			})
			req, _ := http.NewRequest("OPTIONS", "/test", nil)
			r.ServeHTTP(w, req)
			So(w.Code, ShouldEqual, http.StatusNoContent)
		})
	})
}
