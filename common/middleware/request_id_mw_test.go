package middleware

import (
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestIdMW(t *testing.T) {
	Convey("Test RequestIdMW", t, func() {
		w := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		_, r := gin.CreateTestContext(w)
		r.Use(RequestIdMW())
		r.GET("/test", func(c *gin.Context) {
			id, _ := c.Get("request_id")
			So(id.(string), ShouldNotEqual, "")
			c.String(http.StatusOK, "ok")
		})
		req, _ := http.NewRequest("GET", "/test", nil)
		r.ServeHTTP(w, req)
		So(w.Body.String(), ShouldEqual, "ok")
		So(w.Header().Get("X-Request-Id"), ShouldNotEqual, "")
	})
}
