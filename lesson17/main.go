package main

//gorm
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID     int
	Name   string
	Gender string
	Hobby  string
}

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&UserInfo{})
	//db.Create(&UserInfo{Name: "小米", Gender: "女", Hobby: "唱歌"})
	//db.Create(&UserInfo{Name: "张三", Gender: "男", Hobby: "游戏"})
	//var u UserInfo
	//db.Debug().First(&u)
	//fmt.Println(u)
	//var u UserInfo
	//db.Debug().First(&u,2)
	//fmt.Println(u)
	var u UserInfo
	db.Debug().First(&u, "name=?", "张三")
	fmt.Println(u)

	db.Debug().Model(&u).Update("hobby", "游戏")
	fmt.Println(u)

	// Update - 更新多个字段
	db.Model(&u).Updates(UserInfo{Name: "李四", Hobby: "跑步"}) // 仅更新非零值字段

	//Delete - 删除
	var u1 UserInfo
	db.Delete(&u1, 2)

}
