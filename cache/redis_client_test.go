package cache

import (
	"context"
	"errors"
	"gg_web_tmpl/common/config"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/go-redis/redis/v8"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func init() {
	config.InitMockConfig()
	InitMockClient()
}

func TestInitRedis(t *testing.T) {
	Convey("Test InitRedis", t, func() {
		Convey("should success", func() {
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			Mocker.ExpectPing().SetVal("Pong")
			patches.ApplyFuncReturn(redis.NewClient, client)
			err := InitRedis(context.Background())
			So(err, ShouldBeNil)
		})
		Convey("should fail", func() {
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			Mocker.ExpectPing().SetErr(errors.New("test error"))
			patches.ApplyFuncReturn(redis.NewClient, client)
			err := InitRedis(context.Background())
			So(err, ShouldNotBeNil)
		})
	})
}
