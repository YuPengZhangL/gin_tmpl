package config

import (
	"errors"
	"github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"testing"
)

func TestInitConfig(t *testing.T) {
	Convey("Test InitConfig", t, func() {
		Convey("should success", func() {
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			configStr := `app:
  name: gg_web_tmpl
  release: false
  port: 8081
  log_level: debug
  log_path: ./logs
  jwt_sign_key: gg_web_tmpl
  jwt_issuer: gg_web_tmpl

mysql:
  user: root
  password: root
  host: 127.0.0.1
  port: 3306
  db: gg_web_tmpl

redis:
  host: 127.0.0.1
  port: 6379`
			patches.ApplyFuncReturn(ioutil.ReadFile, []byte(configStr), nil)
			cfg := &Config{
				App: &App{
					Name:       "gg_web_tmpl",
					Release:    false,
					Port:       8081,
					LogLevel:   "debug",
					LogPath:    "./logs",
					JWTSignKey: "gg_web_tmpl",
					JWTIssuer:  "gg_web_tmpl",
				},
				MySQL: &MySQL{
					User:     "root",
					Password: "root",
					DB:       "gg_web_tmpl",
					Host:     "127.0.0.1",
					Port:     3306,
				},
				Redis: &Redis{
					RedisHost: "127.0.0.1",
					RedisPort: "6379",
				},
			}
			err := InitConfig("test")
			So(err, ShouldBeNil)
			So(conf, ShouldResemble, cfg)
		})
		Convey("should fail: io error", func() {
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFuncReturn(ioutil.ReadFile, nil, errors.New("test error"))
			err := InitConfig("test")
			So(err, ShouldNotBeNil)
		})
		Convey("should fail: yaml error", func() {
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFuncReturn(ioutil.ReadFile, []byte("}%^&{"), nil)
			err := InitConfig("test")
			So(err, ShouldNotBeNil)
		})
	})
}

func TestInitMockConfig(t *testing.T) {
	Convey("Test InitMockConfig", t, func() {
		cfg := &Config{
			App: &App{
				Name:       "gg_web_tmpl_test",
				Release:    true,
				Port:       8081,
				LogLevel:   "debug",
				LogPath:    "./logs",
				JWTSignKey: "test",
				JWTIssuer:  "test",
			},
			MySQL: &MySQL{
				User:     "test",
				Password: "test",
				DB:       "test",
				Host:     "127.0.0.1",
				Port:     3306,
			},
			Redis: &Redis{
				RedisHost: "127.0.0.1",
				RedisPort: "6379",
			},
		}
		InitMockConfig()
		So(conf, ShouldResemble, cfg)
	})
}
