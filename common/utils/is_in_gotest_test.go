package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsInTest(t *testing.T) {
	Convey("Test IsInTest", t, func() {
		is := IsInTest()
		So(is, ShouldBeTrue)
	})
}
