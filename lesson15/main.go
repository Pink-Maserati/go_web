package main

//路由和路由组
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "get",
		})

	})
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "post",
		})

	})

	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "put",
		})

	})
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "delete",
		})

	})
	//Any可以匹配所有请求方法
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{
				"method": "get",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{
				"method": "post",
			})

		}
		//c.JSON(http.StatusOK,gin.H{
		//	"method":"Any",
		//})

	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "路由不存在",
		})

	})
	//路由组 :多用于区分不同的业务线和API版本
	group := r.Group("/video")
	{
		group.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "video/index",
			})

		})
		group.GET("/haha", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "video/haha",
			})

		})

		group.POST("/shop", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "video/shop",
			})

		})
	}

	r.Run()
}
