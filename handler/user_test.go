package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"gg_web_tmpl/common/consts"
	"gg_web_tmpl/common/log"
	"gg_web_tmpl/common/resp"
	"gg_web_tmpl/service"
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
}

func TestRegisterUser(t *testing.T) {
	Convey("Test RegisterUser", t, func() {
		Convey("should success", func() {
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.POST("/test", RegisterUser)
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFunc(service.RegisterUser, func(c *gin.Context, _ *vo.RegisterUserRequest) error {
				c.String(http.StatusOK, "ok")
				return nil
			})
			regReq := vo.RegisterUserRequest{}
			jsonValue, _ := json.Marshal(&regReq)
			req, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(jsonValue))
			r.ServeHTTP(w, req)
			So(w.Body.String(), ShouldEqual, "ok")
		})
		Convey("should fail: bind json error", func() {
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.POST("/test", RegisterUser)
			req, _ := http.NewRequest("POST", "/test", bytes.NewBuffer([]byte("{1")))
			r.ServeHTTP(w, req)
			rsp := resp.NewBaseResp(consts.CodeParameterError, "parameter error")
			got := &resp.BaseResp{}
			err := json.Unmarshal(w.Body.Bytes(), got)
			So(err, ShouldBeNil)
			So(got, ShouldResemble, rsp)
		})
		Convey("should fail: service error", func() {
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.POST("/test", RegisterUser)
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFuncReturn(service.RegisterUser, errors.New("test error"))
			regReq := vo.RegisterUserRequest{}
			jsonValue, _ := json.Marshal(&regReq)
			req, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(jsonValue))
			r.ServeHTTP(w, req)
			rsp := resp.NewBaseResp(consts.CodeFail, "fail")
			got := &resp.BaseResp{}
			err := json.Unmarshal(w.Body.Bytes(), got)
			So(err, ShouldBeNil)
			So(got, ShouldResemble, rsp)
		})
	})
}

func TestLogin(t *testing.T) {
	Convey("Test Login", t, func() {
		Convey("should success", func() {
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.POST("/test", Login)
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFunc(service.Login, func(c *gin.Context, _ *vo.LoginUserRequest) error {
				c.String(http.StatusOK, "ok")
				return nil
			})
			regReq := vo.LoginUserRequest{}
			jsonValue, _ := json.Marshal(&regReq)
			req, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(jsonValue))
			r.ServeHTTP(w, req)
			So(w.Body.String(), ShouldEqual, "ok")
		})
		Convey("should fail: bind json error", func() {
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.POST("/test", Login)
			req, _ := http.NewRequest("POST", "/test", bytes.NewBuffer([]byte("{1")))
			r.ServeHTTP(w, req)
			rsp := resp.NewBaseResp(consts.CodeParameterError, "parameter error")
			got := &resp.BaseResp{}
			err := json.Unmarshal(w.Body.Bytes(), got)
			So(err, ShouldBeNil)
			So(got, ShouldResemble, rsp)
		})
		Convey("should fail: service error", func() {
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.POST("/test", Login)
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFuncReturn(service.Login, errors.New("test error"))
			regReq := vo.RegisterUserRequest{}
			jsonValue, _ := json.Marshal(&regReq)
			req, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(jsonValue))
			r.ServeHTTP(w, req)
			rsp := resp.NewBaseResp(consts.CodeFail, "fail")
			got := &resp.BaseResp{}
			err := json.Unmarshal(w.Body.Bytes(), got)
			So(err, ShouldBeNil)
			So(got, ShouldResemble, rsp)
		})
	})
}

func TestGetUserInfo(t *testing.T) {
	Convey("Test GetUserInfo", t, func() {
		Convey("should success", func() {
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			// 向context中注入user_id
			r.GET("/test", func(c *gin.Context) {
				c.Set("user_id", "1001")
				GetUserInfo(c)
			})
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFunc(service.GetUserInfo, func(c *gin.Context, _ int64) error {
				c.String(http.StatusOK, "ok")
				return nil
			})
			req, _ := http.NewRequest("GET", "/test", nil)
			r.ServeHTTP(w, req)
			So(w.Body.String(), ShouldEqual, "ok")
		})
		Convey("should fail: user_id not exist", func() {
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			r.GET("/test", GetUserInfo)
			req, _ := http.NewRequest("GET", "/test", nil)
			r.ServeHTTP(w, req)
			rsp := resp.NewBaseResp(consts.CodeFail, "fail")
			got := &resp.BaseResp{}
			err := json.Unmarshal(w.Body.Bytes(), got)
			So(err, ShouldBeNil)
			So(got, ShouldResemble, rsp)
		})
		Convey("should fail: service error", func() {
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			_, r := gin.CreateTestContext(w)
			// 向context中注入user_id
			r.GET("/test", func(c *gin.Context) {
				c.Set("user_id", "1001")
				GetUserInfo(c)
			})
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFuncReturn(service.GetUserInfo, errors.New("test error"))
			req, _ := http.NewRequest("GET", "/test", nil)
			r.ServeHTTP(w, req)
			rsp := resp.NewBaseResp(consts.CodeFail, "fail")
			got := &resp.BaseResp{}
			err := json.Unmarshal(w.Body.Bytes(), got)
			So(err, ShouldBeNil)
			So(got, ShouldResemble, rsp)
		})
	})
}
