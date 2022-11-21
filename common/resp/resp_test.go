package resp

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewBaseResp(t *testing.T) {
	Convey("Test NewBaseResp", t, func() {
		resp := NewBaseResp(0, "ok")
		So(resp.StatusCode, ShouldEqual, 0)
		So(resp.StatusMsg, ShouldEqual, "ok")
		resp = NewBaseResp(1, "test")
		So(resp.StatusCode, ShouldEqual, 1)
		So(resp.StatusMsg, ShouldEqual, "test")
	})
}
