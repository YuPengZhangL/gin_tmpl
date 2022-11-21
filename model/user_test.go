package model

import (
	"context"
	"errors"
	"gg_web_tmpl/common/log"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
	"testing"
)

func init() {
	InitMock()
	log.InitMockLogger()
}

func TestSelectUserByID(t *testing.T) {
	Convey("Test SelectUserByID", t, func() {
		Convey("should success", func() {
			sql := "SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT 1"
			rows := sqlmock.NewRows([]string{"id", "username", "password"})
			rows.AddRow(1, "user", "password")
			Mocker.ExpectQuery(sql).WithArgs(1).WillReturnRows(rows)
			user := &User{
				ID:       1,
				Username: "user",
				Password: "password",
			}
			got, err := SelectUserByID(context.Background(), GetDB(), 1)
			So(err, ShouldBeNil)
			So(got, ShouldResemble, user)
		})
		Convey("should success: record not found", func() {
			sql := "SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT 1"
			Mocker.ExpectQuery(sql).WithArgs(1).WillReturnError(gorm.ErrRecordNotFound)
			got, err := SelectUserByID(context.Background(), GetDB(), 1)
			So(err, ShouldBeNil)
			So(got, ShouldBeNil)
		})
		Convey("should fail", func() {
			sql := "SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT 1"
			Mocker.ExpectQuery(sql).WithArgs(1).WillReturnError(errors.New("test error"))
			got, err := SelectUserByID(context.Background(), GetDB(), 1)
			So(err, ShouldNotBeNil)
			So(got, ShouldBeNil)
		})
	})
}

func TestSelectUserByUsername(t *testing.T) {
	Convey("Test SelectUserByUsername", t, func() {
		Convey("should success", func() {
			rows := sqlmock.NewRows([]string{"id", "username", "password"})
			rows.AddRow(1, "user", "password")
			sql := "SELECT * FROM `users` WHERE username = ? ORDER BY `users`.`id` LIMIT 1"
			Mocker.ExpectQuery(sql).WithArgs("user").WillReturnRows(rows)
			user := &User{
				ID:       1,
				Username: "user",
				Password: "password",
			}
			got, err := SelectUserByUsername(context.Background(), GetDB(), "user")
			So(err, ShouldBeNil)
			So(got, ShouldResemble, user)
		})
		Convey("should success: record not found", func() {
			sql := "SELECT * FROM `users` WHERE username = ? ORDER BY `users`.`id` LIMIT 1"
			Mocker.ExpectQuery(sql).WithArgs("user").WillReturnError(gorm.ErrRecordNotFound)
			got, err := SelectUserByUsername(context.Background(), GetDB(), "user")
			So(err, ShouldBeNil)
			So(got, ShouldBeNil)
		})
		Convey("should fail", func() {
			sql := "SELECT * FROM `users` WHERE username = ? ORDER BY `users`.`id` LIMIT 1"
			Mocker.ExpectQuery(sql).WithArgs("user").WillReturnError(errors.New("test error"))
			got, err := SelectUserByUsername(context.Background(), GetDB(), "user")
			So(err, ShouldNotBeNil)
			So(got, ShouldBeNil)
		})
	})
}

func TestInsertUser(t *testing.T) {
	Convey("Test InsertUser", t, func() {
		Convey("should success", func() {
			user := &User{
				ID:       1,
				Username: "user",
				Password: "password",
			}
			Mocker.ExpectBegin()
			Mocker.ExpectExec("INSERT INTO `users` (`username`,`password`,`id`) VALUES (?,?,?)").WithArgs("user", "password", 1).WillReturnResult(sqlmock.NewResult(1, 1))
			Mocker.ExpectCommit()
			got, err := InsertUser(context.Background(), GetDB(), user)
			So(err, ShouldBeNil)
			So(got, ShouldResemble, got)
		})
	})
}
