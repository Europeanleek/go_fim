package core

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//	func InitMysql() *gorm.DB {
//		dsn := "root:12345678@tcp(127.0.0.1:33069)/fim_server_db?charset=utf8mb4&parseTime=True&loc=Local"
//		var mysqlLogger logger.Interface
//		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
//			Logger: mysqlLogger,
//		})
//		if err != nil {
//			log.Fatalf(fmt.Sprintf("[%s] mysql连接失败", dsn))
//		}
//		sqlDB, _ := db.DB()
//		sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
//		sqlDB.SetMaxOpenConns(100)              // 最多可容纳
//		sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间，不能超过mysql的wait_timeout
//		log.Default().Println("mysql连接成功")
//		return db
//	}
func InitGorm(MysqlDataSource string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(MysqlDataSource), &gorm.Config{})
	if err != nil {
		panic("连接mysql数据库失败, error=" + err.Error())
	} else {
		fmt.Println("连接mysql数据库成功")
	}
	return db
}
