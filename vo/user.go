package vo

import (
	"gg_web_tmpl/model"
	"github.com/spf13/cast"
)

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterUserResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	Data       struct {
		Token  string `json:"token"`   // 用户鉴权token
		UserID string `json:"user_id"` // 用户id
	} `json:"data"`
}

type GetUserInfoResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	Data       struct {
		UserID   string `json:"user_id"` // 用户id
		UserName string `json:"user_name"`
	} `json:"data"`
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	Data       struct {
		Token  string `json:"token"`   // 用户鉴权token
		UserID string `json:"user_id"` // 用户id
	} `json:"data"`
}

func NewRegisterUserResponse(token string, userID int64) *RegisterUserResponse {
	return &RegisterUserResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		Data: struct {
			Token  string `json:"token"`   // 用户鉴权token
			UserID string `json:"user_id"` // 用户id
		}{
			Token:  token,
			UserID: cast.ToString(userID),
		},
	}
}

func NewLoginResponse(token string, userID int64) *LoginResponse {
	return &LoginResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		Data: struct {
			Token  string `json:"token"`   // 用户鉴权token
			UserID string `json:"user_id"` // 用户id
		}{
			Token:  token,
			UserID: cast.ToString(userID),
		},
	}
}

func NewGetUserInfoResponse(user *model.User) *GetUserInfoResponse {
	return &GetUserInfoResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		Data: struct {
			UserID   string `json:"user_id"`   // 用户id
			UserName string `json:"user_name"` // 用户名
		}{
			UserID:   cast.ToString(user.ID),
			UserName: user.Username,
		},
	}
}
