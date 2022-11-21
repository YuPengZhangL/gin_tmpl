package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCompareHashAndPassword(t *testing.T) {
	Convey("Test CompareHashAndPassword", t, func() {
		before := "password"
		after := GenerateHashFromPassword(before)
		So(after, ShouldNotEqual, "")
		ok := CompareHashAndPassword(after, "password")
		So(ok, ShouldBeTrue)
		ok = CompareHashAndPassword(after, "test")
		So(ok, ShouldBeFalse)
	})
}
