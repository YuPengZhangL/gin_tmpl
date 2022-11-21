package model

import (
	"errors"
	"gg_web_tmpl/common/config"
	"gg_web_tmpl/common/log"
	"github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
	"testing"
)

func init() {
	InitMock()
	log.InitMockLogger()
	config.InitMockConfig()
}

func TestInitMySQL(t *testing.T) {
	Convey("Test InitMySQL", t, func() {
		Convey("should success", func() {
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFuncReturn(gorm.Open, db, nil)
			err := InitMySQL(config.GetConf().MySQL)
			So(err, ShouldBeNil)
		})
		Convey("should fail", func() {
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			patches.ApplyFuncReturn(gorm.Open, nil, errors.New("test error"))
			err := InitMySQL(config.GetConf().MySQL)
			So(err, ShouldNotBeNil)
		})
	})
}
