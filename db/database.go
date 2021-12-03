package db

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DBConn struct {
	DB *gorm.DB
}

func GetDB() *gorm.DB {
	return DB.DB
}

var DB *DBConn

func InitDB() {
	// 数据库连接

	dsn := "root:123456@tcp(localhost:3306)/gin-web?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 打印所有执行的SQL
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名不加s
		},
	})
	if err != nil {

		print(err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		_ = sqlDB.Close()
		print(err.Error())
	}
	sqlDB.SetMaxIdleConns(100) // 设置最大连接数
	sqlDB.SetMaxOpenConns(50) // 设置最大空闲连接数

	// 获取sql配置情况
	sqlStats, _ := json.Marshal(sqlDB.Stats())
	fmt.Println(string(sqlStats))
	DB = &DBConn{
		DB: db,
	}

}

func (db DBConn) Close() {
	sqlDB, err := db.DB.DB()
	if err != nil {
		_ = sqlDB.Close()
		print(err.Error())
	}
	sqlDB.Close()
}
