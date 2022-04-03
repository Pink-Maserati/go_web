package main

//获取form表单提交的参数
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("login.html", "index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		//username:=c.PostForm("username")
		//password:=c.PostForm("password")
		//username := c.DefaultPostForm("username", "somebody")
		//password := c.DefaultPostForm("xxx", "***")
		username, ok := c.GetPostForm("username")
		if !ok {
			username = "sb"
		}
		password, _ := c.GetPostForm("password")

		c.HTML(http.StatusOK, "index.html", gin.H{
			"username": username,
			"password": password,
		})
	})

	r.Run(":9090")

}
