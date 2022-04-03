package main

import (
	"fmt"
	"go_web/todo/dao"
	"go_web/todo/models"
	"go_web/todo/routers"
	"go_web/todo/setting"
	"os"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fast := ListNode{0, nil}
	fmt.Printf("%T %v %v\n", fast, fast, fast.Val)

	return

	if len(os.Args) < 2 {
		fmt.Println("Usage：./todo conf/config.ini")
		return
	}

	//加载配置文件
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config file failed,err:%v\n", err)
		return
	}

	fmt.Printf("config:%v\n", setting.Conf)

	//数据库初始化
	err := dao.InitMysql(setting.Conf.MySQLConfig)
	if err != nil {
		panic(err)
	}

	err = dao.DB.AutoMigrate(&models.Todo{})
	if err != nil {
		panic(err)
	}

	r := routers.SetupRouter(setting.Conf.LoggerConfig)

	err = r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("server start failed,err:%v\n", err)
	}

}
