package model

import (
	"fmt"
	"gg_web_tmpl/common/config"
	"gg_web_tmpl/common/log"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	db     *gorm.DB
	Mocker sqlmock.Sqlmock
)

func InitMock() {
	mysqlDB, m, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	Mocker = m
	db, _ = gorm.Open(mysql.New(mysql.Config{
		Conn:                      mysqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: log.NewGORMLogger(),
	})
}

// InitMySQL 初始化MySQL
func InitMySQL(cfg *config.MySQL) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: log.NewGORMLogger(),
	})
	if err != nil {
		return err
	}
	db = gormDB
	if err = db.AutoMigrate(&User{}); err != nil {
		return err
	}
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return nil
}

func GetDB() *gorm.DB {
	return db
}
