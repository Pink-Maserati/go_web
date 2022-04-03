package main

//重定向和转发
//重定向地址会变，转发地址不会
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//谷歌调转是http://127.0.0.1:9090/baidu.com 加了本地的地址和端口 火狐测试正常
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
		//c.JSON(http.StatusOK,gin.H{
		//	"status":"ok",
		//})
	})
	//转发
	r.GET("/a", func(c *gin.Context) {
		//跳转到/b对应的路由处理函数
		c.Request.URL.Path = "/b" //把请求中的URI修改
		r.HandleContext(c)        //继续后续的处理

	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})

	})
	r.Run(":9090")

}
