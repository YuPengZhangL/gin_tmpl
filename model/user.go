package model

import (
	"context"
	"gg_web_tmpl/common/log"
	"gorm.io/gorm"
)

type User struct {
	ID       int64  `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

func (User) TableName() string {
	return "users"
}

func SelectUserByID(ctx context.Context, db *gorm.DB, id int64) (*User, error) {
	user := &User{}
	err := db.Model(&User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.GetLoggerWithCtx(ctx).Errorf("select user by id err:%v", err)
		return nil, err
	}
	return user, nil
}

func SelectUserByUsername(ctx context.Context, db *gorm.DB, username string) (*User, error) {
	user := &User{}
	err := db.Model(&User{}).Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.GetLoggerWithCtx(ctx).Errorf("select user by username err:%v", err)
		return nil, err
	}
	return user, nil
}

func InsertUser(ctx context.Context, db *gorm.DB, user *User) (*User, error) {
	err := db.Model(&User{}).Create(user).Error
	if err != nil {
		log.GetLoggerWithCtx(ctx).Errorf("insert user err:%v", err)
		return nil, err
	}
	return user, nil
}
