package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello goLang!",
	})

}
func main2() {
	//创建一个默认的路由引擎
	r := gin.Default()
	r.GET("./hello", hello)
	r.GET("./book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "get",
		})

	})
	//restful
	r.POST("./book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "post",
		})

	})
	r.PUT("./book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "put",
		})

	})
	r.DELETE("./book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "delete",
		})

	})
	//启动HTTP服务
	r.Run(":9090")

}
