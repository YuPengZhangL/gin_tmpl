package service

import (
	"errors"
	"gg_web_tmpl/cache"
	"gg_web_tmpl/common/config"
	"gg_web_tmpl/common/log"
	"gg_web_tmpl/common/utils"
	"gg_web_tmpl/model"
	"gg_web_tmpl/vo"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	log.InitMockLogger()
	config.InitMockConfig()
	cache.InitMockClient()
}

func TestRegisterUser(t *testing.T) {
	Convey("Test RegisterUser", t, func() {
		Convey("should success", func() {
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFuncReturn(model.InsertUser, &model.User{ID: 1}, nil)
			patches.ApplyFuncReturn(utils.GenerateToken, "token", nil)
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.GET("/test", func(c *gin.Context) {
				err := RegisterUser(c, &vo.RegisterUserRequest{
					Username: "test",
					Password: "test",
				})
				So(err, ShouldBeNil)
			})

			req, _ := http.NewRequest("GET", "/test", nil)
			r.ServeHTTP(w, req)
			So(w.Body.String(), ShouldEqual, `{"status_code":0,"status_msg":"ok","data":{"token":"token","user_id":"1"}}`)
		})
		Convey("should fail: insert error", func() {
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFuncReturn(model.InsertUser, nil, errors.New("test error"))
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.GET("/test", func(c *gin.Context) {
				err := RegisterUser(c, &vo.RegisterUserRequest{
					Username: "test",
					Password: "test",
				})
				So(err, ShouldNotBeNil)
			})

			req, _ := http.NewRequest("GET", "/test", nil)
			r.ServeHTTP(w, req)
		})
		Convey("should fail: token error", func() {
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFuncReturn(model.InsertUser, &model.User{ID: 1}, nil)
			patches.ApplyFuncReturn(utils.GenerateToken, nil, errors.New("test error"))
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.GET("/test", func(c *gin.Context) {
				err := RegisterUser(c, &vo.RegisterUserRequest{
					Username: "test",
					Password: "test",
				})
				So(err, ShouldNotBeNil)
			})

			req, _ := http.NewRequest("GET", "/test", nil)
			r.ServeHTTP(w, req)
		})
	})
}

func TestLogin(t *testing.T) {
	Convey("Test Login", t, func() {
		Convey("should success: normal", func() {
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFuncReturn(model.SelectUserByUsername, &model.User{ID: 1}, nil)
			patches.ApplyFuncReturn(utils.CompareHashAndPassword, true)
			patches.ApplyFuncReturn(utils.GenerateToken, "token", nil)
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.GET("/test", func(c *gin.Context) {
				err := Login(c, &vo.LoginUserRequest{
					Username: "test",
					Password: "test",
				})
				So(err, ShouldBeNil)
			})

			req, _ := http.NewRequest("GET", "/test", nil)
			r.ServeHTTP(w, req)
			So(w.Body.String(), ShouldEqual, `{"status_code":0,"status_msg":"ok","data":{"token":"token","user_id":"1"}}`)
		})
		Convey("should success: login fail, user not exist", func() {
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFuncReturn(model.SelectUserByUsername, nil, nil)
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.GET("/test", func(c *gin.Context) {
				err := Login(c, &vo.LoginUserRequest{
					Username: "test",
					Password: "test",
				})
				So(err, ShouldBeNil)
			})

			req, _ := http.NewRequest("GET", "/test", nil)
			r.ServeHTTP(w, req)
			So(w.Body.String(), ShouldEqual, `{"status_code":-1,"status_msg":"login fail"}`)
		})
		Convey("should success: login fail, password incorrect", func() {
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFuncReturn(model.SelectUserByUsername, &model.User{ID: 1}, nil)
			patches.ApplyFuncReturn(utils.CompareHashAndPassword, false)
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.GET("/test", func(c *gin.Context) {
				err := Login(c, &vo.LoginUserRequest{
					Username: "test",
					Password: "test",
				})
				So(err, ShouldBeNil)
			})

			req, _ := http.NewRequest("GET", "/test", nil)
			r.ServeHTTP(w, req)
			So(w.Body.String(), ShouldEqual, `{"status_code":-1,"status_msg":"login fail"}`)
		})
		Convey("should fail: token error", func() {
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFuncReturn(model.SelectUserByUsername, &model.User{ID: 1}, nil)
			patches.ApplyFuncReturn(utils.CompareHashAndPassword, true)
			patches.ApplyFuncReturn(utils.GenerateToken, "", errors.New("test error"))
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.GET("/test", func(c *gin.Context) {
				err := Login(c, &vo.LoginUserRequest{
					Username: "test",
					Password: "test",
				})
				So(err, ShouldNotBeNil)
			})

			req, _ := http.NewRequest("GET", "/test", nil)
			r.ServeHTTP(w, req)
		})
	})
}

func TestGetUserInfo(t *testing.T) {
	Convey("Test GetUserInfo", t, func() {
		Convey("should success: cache exist", func() {
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFuncReturn(cache.GetUser, &model.User{ID: 1}, nil)
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.GET("/test", func(c *gin.Context) {
				err := GetUserInfo(c, 1)
				So(err, ShouldBeNil)
			})

			req, _ := http.NewRequest("GET", "/test", nil)
			r.ServeHTTP(w, req)
			So(w.Body.String(), ShouldEqual, `{"status_code":0,"status_msg":"ok","data":{"user_id":"1","user_name":""}}`)
		})
		Convey("should success: cache not exist, get from db", func() {
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFuncReturn(cache.GetUser, nil, nil)
			patches.ApplyFuncReturn(model.SelectUserByID, &model.User{ID: 1}, nil)
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.GET("/test", func(c *gin.Context) {
				err := GetUserInfo(c, 1)
				So(err, ShouldBeNil)
			})

			req, _ := http.NewRequest("GET", "/test", nil)
			r.ServeHTTP(w, req)
			So(w.Body.String(), ShouldEqual, `{"status_code":0,"status_msg":"ok","data":{"user_id":"1","user_name":""}}`)
		})
		Convey("should success: cache not exist, db not exist", func() {
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFuncReturn(cache.GetUser, nil, nil)
			patches.ApplyFuncReturn(model.SelectUserByID, nil, nil)
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.GET("/test", func(c *gin.Context) {
				err := GetUserInfo(c, 1)
				So(err, ShouldBeNil)
			})

			req, _ := http.NewRequest("GET", "/test", nil)
			r.ServeHTTP(w, req)
			So(w.Body.String(), ShouldEqual, `{"status_code":-1,"status_msg":"user not exist"}`)
		})
		Convey("should fail: db error", func() {
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFuncReturn(cache.GetUser, nil, nil)
			patches.ApplyFuncReturn(model.SelectUserByID, nil, errors.New("test error"))
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.GET("/test", func(c *gin.Context) {
				err := GetUserInfo(c, 1)
				So(err, ShouldNotBeNil)
			})

			req, _ := http.NewRequest("GET", "/test", nil)
			r.ServeHTTP(w, req)
		})
	})
}
