package routers

import (
	"github.com/gin-gonic/gin"
	"go_web/todo/controller"
	middleware "go_web/todo/logger"
	"go_web/todo/setting"
	"net/http"
)

func SetupRouter(loggerConfig *setting.LoggerConfig) *gin.Engine {

	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	//加载静态文件
	r.Static("static", "./static")

	//模板解析
	r.LoadHTMLGlob("templates/*")

	//中间件
	r.Use(middleware.LoggerToFile(loggerConfig))

	//路由注册
	r.GET("/", controller.IndexHandler)

	// 中间件 统计耗时+ 日志请求

	// v1 组
	v1Group := r.Group("v1")
	{
		//添加事项
		v1Group.POST("/todo", controller.CreateTodo)

		//查看所有的待办事项
		v1Group.GET("/todo", controller.GetAllTodos)

		//删除某一个代办事项
		v1Group.DELETE("/todo/:id", controller.DeleteTodoById)

		//更新、修改某一个代办事项
		v1Group.PUT("/todo/:id", controller.UpdateTodoById)

	}

	//当请求的路由不存在时，默认返回
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "路由不存在",
		})

	})

	return r
}
