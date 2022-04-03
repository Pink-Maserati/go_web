package main

//中间件
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	fmt.Println("index")
	//从上下文中取值（跨中间件取值）
	name, ok := c.Get("name")
	if !ok {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

//中间件： 统计请求处理函数耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	//计时
	start := time.Now()

	// go funcXX(c.copy()) 在funcXX中只能使用c的拷贝

	//调用该请求的剩余处理程序
	c.Next()
	//不调用该请求的剩余处理程序
	//c.Abort()
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out...")

}

func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	//c.Next() //调用后续的处理函数
	//c.Abort() //阻止调用后续的处理函数
	//return
	//在上下文c中设置值
	c.Set("name", "Go")
	fmt.Println("m2 out...")

}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	//连接数据库
	//或者做一些其他的准备工作
	return func(c *gin.Context) {
		if doCheck {
			//存放具体的逻辑
			c.Next()
		} else {
			c.Next()
		}

	}

}

func main() {

	//r :=gin.Default()//默认使用Logger(), Recovery
	r := gin.New()
	//全局注册中间件
	r.Use(m1, m2, authMiddleware(false))
	//r.GET("/index", m1, indexHandler)
	//r.GET("/shop",m1, func(c *gin.Context) {
	//	c.JSON(http.StatusOK,gin.H{
	//		"msg":"shop",
	//	})
	//})
	r.GET("/index", indexHandler)
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})

	//为路由组注册中间件方法1：
	shopGroup := r.Group("/shop", authMiddleware(true))
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "shop/index",
			})
		})
	}

	//为路由组注册中间件方法2：
	shopGroup2 := r.Group("/shop2")
	shopGroup2.Use(authMiddleware(true))
	{
		shopGroup2.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "shop2/index",
			})
		})
	}
	r.Run()

}
