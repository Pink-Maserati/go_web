package dao

import (
	"fmt"
	"go_web/todo/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

//数据库初始化
func InitMysql(cfg *setting.MySQLConfig) (err error) {
	//创建数据库sql create database todo;
	//数据库连接
	//dsn := "root:password@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=true&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	return
}
