package cache

import (
	"context"
	"encoding/json"
	"errors"
	"gg_web_tmpl/common/log"
	"gg_web_tmpl/model"
	"github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func init() {
	InitMockClient()
	log.InitMockLogger()
}

func TestSetUser(t *testing.T) {
	Convey("Test SetUser", t, func() {
		Convey("should success", func() {
			defer Mocker.ClearExpect()
			user := &model.User{
				ID:       1001,
				Username: "test",
				Password: "test",
			}
			str, _ := json.Marshal(user)
			Mocker.ExpectSet(getUserKey(1001), str, 0).SetVal("")
			err := SetUser(context.Background(), GetRedisClient(), user)
			So(err, ShouldBeNil)
		})
		Convey("should fail: json error", func() {
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFuncReturn(json.Marshal, nil, errors.New("test error"))
			err := SetUser(context.Background(), GetRedisClient(), &model.User{})
			So(err, ShouldNotBeNil)
		})
		Convey("should fail: redis error", func() {
			defer Mocker.ClearExpect()
			user := &model.User{
				ID:       1001,
				Username: "test",
				Password: "test",
			}
			str, _ := json.Marshal(user)
			Mocker.ExpectSet(getUserKey(1001), string(str), 0).SetErr(errors.New("test error"))
			err := SetUser(context.Background(), GetRedisClient(), user)
			So(err, ShouldNotBeNil)
		})
	})
}

func TestGetUser(t *testing.T) {
	Convey("Test GetUser", t, func() {
		Convey("should success: normal", func() {
			defer Mocker.ClearExpect()
			user := &model.User{
				ID:       1001,
				Username: "test",
				Password: "test",
			}
			str, _ := json.Marshal(user)
			Mocker.ExpectGet(getUserKey(1001)).SetVal(string(str))
			got, err := GetUser(context.Background(), GetRedisClient(), 1001)
			So(err, ShouldBeNil)
			So(got, ShouldNotBeNil)
			So(got, ShouldResemble, user)
		})
		Convey("should success: redis.Nil", func() {
			defer Mocker.ClearExpect()
			Mocker.ExpectGet(getUserKey(1001)).RedisNil()
			got, err := GetUser(context.Background(), GetRedisClient(), 1001)
			So(err, ShouldBeNil)
			So(got, ShouldBeNil)
		})
		Convey("should fail: json error", func() {
			defer Mocker.ClearExpect()
			Mocker.ExpectGet(getUserKey(1001)).SetVal("{1")
			got, err := GetUser(context.Background(), GetRedisClient(), 1001)
			So(err, ShouldNotBeNil)
			So(got, ShouldBeNil)
		})
		Convey("should fail: redis error", func() {
			defer Mocker.ClearExpect()
			Mocker.ExpectGet(getUserKey(1001)).SetErr(errors.New("test error"))
			got, err := GetUser(context.Background(), GetRedisClient(), 1001)
			So(err, ShouldNotBeNil)
			So(got, ShouldBeNil)
		})

	})
}
