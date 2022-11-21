package handler

import (
	"gg_web_tmpl/common/consts"
	"gg_web_tmpl/common/log"
	"gg_web_tmpl/common/resp"
	"gg_web_tmpl/service"
	"gg_web_tmpl/vo"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

// RegisterUser 用户注册
func RegisterUser(c *gin.Context) {
	logger := log.GetLoggerWithCtx(c)
	req := &vo.RegisterUserRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		logger.Errorf("register user parameter error: %v", err)
		c.JSON(http.StatusOK, resp.NewBaseResp(consts.CodeParameterError, "parameter error"))
		return
	}
	err = service.RegisterUser(c, req)
	if err != nil {
		logger.Errorf("register user error: %v", err)
		c.JSON(http.StatusOK, resp.NewBaseResp(consts.CodeFail, err.Error()))
		return
	}
	return
}

// Login 用户登陆
func Login(c *gin.Context) {
	logger := log.GetLoggerWithCtx(c)
	req := &vo.LoginUserRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		logger.Errorf("Login user parameter error: %v", err)
		c.JSON(http.StatusOK, resp.NewBaseResp(consts.CodeParameterError, "parameter error"))
		return
	}
	err = service.Login(c, req)
	if err != nil {
		logger.Errorf("Login user error: %v", err)
		c.JSON(http.StatusOK, resp.NewBaseResp(consts.CodeFail, "fail"))
		return
	}
	return
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	logger := log.GetLoggerWithCtx(c)
	idStr, ok := c.Get("user_id")
	if !ok {
		logger.Errorf("get user_id from context error: not exist")
		c.JSON(http.StatusOK, resp.NewBaseResp(consts.CodeFail, "fail"))
		return
	}
	userID := cast.ToInt64(idStr)
	err := service.GetUserInfo(c, userID)
	if err != nil {
		logger.Errorf("get user info error: %v", userID)
		c.JSON(http.StatusOK, resp.NewBaseResp(consts.CodeFail, "fail"))
		return
	}
	return
}
