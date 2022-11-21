package log

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestInitLogger(t *testing.T) {
	Convey("Test InitLogger", t, func() {
		Convey("should success", func() {
			err := InitLogger("./logs", "debug")
			So(err, ShouldBeNil)
		})
		Convey("should fail", func() {
			err := InitLogger("./logs", "test")
			So(err, ShouldNotBeNil)
		})
	})
}

func TestGetLoggerWithCtx(t *testing.T) {
	Convey("Test GetLoggerWithCtx", t, func() {
		Convey("request_id exist", func() {
			err := InitLogger("./logs", "debug")
			So(err, ShouldBeNil)
			ctx := context.WithValue(context.Background(), "request_id", "test")
			l := GetLoggerWithCtx(ctx)
			id, ok := l.Data["request_id"]
			So(ok, ShouldBeTrue)
			So(id, ShouldEqual, "test")
		})
		Convey("request_id not exist", func() {
			err := InitLogger("./logs", "debug")
			So(err, ShouldBeNil)
			l := GetLoggerWithCtx(context.Background())
			id, ok := l.Data["request_id"]
			So(ok, ShouldBeTrue)
			So(id, ShouldEqual, "")
		})
	})
}
