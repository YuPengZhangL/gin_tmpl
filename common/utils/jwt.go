package utils

import (
	"errors"
	"gg_web_tmpl/common/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cast"
	"time"
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

// CustomClaims 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	jwt.StandardClaims
}

// NewJWT 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(config.GetConf().App.JWTSignKey),
	}
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string, expiresAt time.Time) (string, error) {
	claims, err := j.ParseToken(tokenString)
	if err != nil {
		return "", err
	}
	claims.StandardClaims.ExpiresAt = expiresAt.Unix()
	return j.CreateToken(*claims)
}

// GenerateToken 生成token
func GenerateToken(id int64) (token string, err error) {
	j := NewJWT()
	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        cast.ToString(id),
			NotBefore: time.Now().Unix() - 1000,                  // 签名生效时间
			ExpiresAt: time.Now().Add(time.Hour * 24 * 1).Unix(), // 过期时间 1天
			Issuer:    config.GetConf().App.JWTIssuer,            //签名的发行者
		},
	}

	token, err = j.CreateToken(claims)
	return
}
