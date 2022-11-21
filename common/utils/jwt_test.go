package utils

import (
	"gg_web_tmpl/common/config"
	"github.com/dgrijalva/jwt-go"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func init() {
	config.InitMockConfig()
}

func TestJWT(t *testing.T) {
	Convey("Test JWT", t, func() {
		Convey("test CreateToken", func() {
			j := NewJWT()
			claims := CustomClaims{
				StandardClaims: jwt.StandardClaims{
					Id:        "10001",
					NotBefore: time.Now().Unix() - 1000,
					ExpiresAt: time.Now().Add(time.Hour).Unix(),
					Issuer:    config.GetConf().App.JWTIssuer,
				},
			}
			token, err := j.CreateToken(claims)
			So(err, ShouldBeNil)
			So(token, ShouldNotEqual, "")
		})
		Convey("test ParseToken", func() {
			Convey("should success", func() {
				j := NewJWT()
				claims := CustomClaims{
					StandardClaims: jwt.StandardClaims{
						Id:        "10001",
						NotBefore: time.Now().Unix() - 1000,
						ExpiresAt: time.Now().Add(time.Hour).Unix(),
						Issuer:    config.GetConf().App.JWTIssuer,
					},
				}
				token, err := j.CreateToken(claims)
				So(err, ShouldBeNil)
				So(token, ShouldNotEqual, "")
				got, err := j.ParseToken(token)
				So(err, ShouldBeNil)
				So(got, ShouldResemble, &claims)
			})
			Convey("should fail: token is expired", func() {
				j := NewJWT()
				claims := CustomClaims{
					StandardClaims: jwt.StandardClaims{
						Id:        "10001",
						NotBefore: time.Now().Unix() - 100000,
						ExpiresAt: time.Now().Unix() - 99999,
						Issuer:    config.GetConf().App.JWTIssuer,
					},
				}
				token, err := j.CreateToken(claims)
				So(err, ShouldBeNil)
				So(token, ShouldNotEqual, "")
				got, err := j.ParseToken(token)
				So(err, ShouldNotBeNil)
				So(err, ShouldResemble, TokenExpired)
				So(got, ShouldBeNil)
			})
			Convey("should fail: not a token", func() {
				j := NewJWT()
				got, err := j.ParseToken("test")
				So(err, ShouldNotBeNil)
				So(err, ShouldResemble, TokenMalformed)
				So(got, ShouldBeNil)
			})
			Convey("should fail: token invalid", func() {
				j := NewJWT()
				token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg2MzU1OTQsImp0aSI6IjYiLCJpc3MiOiJ0ZXN0IiwibmJmIjoxNjU4MDI5Nzk0fQ==.H9lJ9cF8A7YRqsmHjko_OqAtW6ux3F4VWnkKg0dfog8`
				got, err := j.ParseToken(token)
				So(err, ShouldNotBeNil)
				So(err, ShouldResemble, TokenInvalid)
				So(got, ShouldBeNil)
			})
		})
		Convey("test RefreshToken", func() {
			Convey("should success", func() {
				j := NewJWT()
				now := time.Now()
				claims := CustomClaims{
					StandardClaims: jwt.StandardClaims{
						Id:        "10001",
						NotBefore: now.Unix() - 1000,
						ExpiresAt: now.Add(time.Minute).Unix(),
						Issuer:    config.GetConf().App.JWTIssuer,
					},
				}
				token, err := j.CreateToken(claims)
				So(err, ShouldBeNil)
				got, err := j.RefreshToken(token, time.Now().Add(time.Hour))
				So(err, ShouldBeNil)
				c, err := j.ParseToken(got)
				So(err, ShouldBeNil)
				So(c.ExpiresAt, ShouldEqual, now.Add(time.Hour).Unix())
			})
			Convey("should fail", func() {
				j := NewJWT()
				got, err := j.RefreshToken("test", time.Now().Add(time.Hour))
				So(err, ShouldNotBeNil)
				So(got, ShouldEqual, "")
			})
		})
	})
}

func TestGenerateToken(t *testing.T) {
	Convey("Test GenerateToken", t, func() {
		token, err := GenerateToken(10001)
		So(err, ShouldBeNil)
		So(token, ShouldNotEqual, "")
		j := NewJWT()
		claims, err := j.ParseToken(token)
		So(err, ShouldBeNil)
		So(claims.Id, ShouldEqual, "10001")
	})
}
