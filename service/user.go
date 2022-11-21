package service

import (
	"fmt"
	"gg_web_tmpl/cache"
	"gg_web_tmpl/common/consts"
	"gg_web_tmpl/common/log"
	"gg_web_tmpl/common/resp"
	"gg_web_tmpl/common/utils"
	"gg_web_tmpl/model"
	"gg_web_tmpl/vo"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterUser 注册
func RegisterUser(c *gin.Context, req *vo.RegisterUserRequest) error {
	logger := log.GetLoggerWithCtx(c)
	pwdHash := utils.GenerateHashFromPassword(req.Password)
	userUsed, err := model.SelectUserByUsername(c, model.GetDB(), req.Username)
	if userUsed != nil {
		return fmt.Errorf("RegisterUser err: 该用户已存在")
	}
	user, err := model.InsertUser(c, model.GetDB(), &model.User{
		Username: req.Username,
		Password: pwdHash,
	})
	if err != nil {
		return fmt.Errorf("RegisterUser err: %v", err)
	}
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return fmt.Errorf("RegisterUser err: %v", err)
	}
	c.JSON(http.StatusOK, vo.NewRegisterUserResponse(token, user.ID))
	logger.WithField("user_id", user.ID).Infof("register user success")
	return nil
}

// Login 登录
func Login(c *gin.Context, req *vo.LoginUserRequest) error {
	logger := log.GetLoggerWithCtx(c)
	user, err := model.SelectUserByUsername(c, model.GetDB(), req.Username)
	if err != nil {
		return fmt.Errorf("RegisterUser err: %v", err)
	}
	if user == nil {
		c.JSON(http.StatusOK, resp.NewBaseResp(consts.CodeFail, "login fail"))
		return nil
	}
	isPwdOk := utils.CompareHashAndPassword(user.Password, req.Password)
	if !isPwdOk {
		c.JSON(http.StatusOK, resp.NewBaseResp(consts.CodeFail, "login fail"))
		return nil
	}
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return fmt.Errorf("RegisterUser err: %v", err)
	}
	c.JSON(http.StatusOK, vo.NewLoginResponse(token, user.ID))
	logger.WithField("user_id", user.ID).Infof("login success")
	return nil
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context, id int64) error {
	logger := log.GetLoggerWithCtx(c).WithField("user_id", id)
	user, err := cache.GetUser(c, cache.GetRedisClient(), id)
	if err != nil {
		logger.Warnf("get user from redis error: %v", err)
	}
	// 缓存不存在
	if user == nil {
		user, err = model.SelectUserByID(c, model.GetDB(), id)
		if err != nil {
			return fmt.Errorf("get user from db err: %v", err)
		}
		// db不存在
		if user == nil {
			c.JSON(http.StatusOK, resp.NewBaseResp(consts.CodeFail, "user not exist"))
			return nil
		}
		// 缓存失败打日志
		err = cache.SetUser(c, cache.GetRedisClient(), user)
		if err != nil {
			logger.Warnf("set user from redis error: %v", err)
		}
	}
	c.JSON(http.StatusOK, vo.NewGetUserInfoResponse(user))
	return nil
}
