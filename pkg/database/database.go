package database

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"

	gormlogger "gorm.io/gorm/logger"
)

// DB 对象
var DB *gorm.DB
var SQLDB *sql.DB

// 连接到数据库

func Connect(dbConfig gorm.Dialector, _logger gormlogger.Interface) {
	var err error

	//还用gorm.Open连接数据库
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	// 获取底层的 sqlDB
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}
