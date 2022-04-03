package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//静态文件：html页面上用到的样式文件.css js文件 图片
func main() {
	r := gin.Default()
	//加载静态文件
	r.Static("statics", "./statics")
	//gin框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//模板解析
	//r.LoadHTMLFiles("templates/index.tmpl")

	r.LoadHTMLGlob("templates/**/*")
	r.GET("/post/index", func(c *gin.Context) {
		//http请求
		c.HTML(http.StatusOK, "posts/index.gotmpl", gin.H{ //模板渲染
			//"title":"/post/index",
			"title": "<a href='https://liwenzhou.com'>李文周的博客</a>",
		})

	})
	r.GET("/user/index", func(c *gin.Context) {
		//http请求
		c.HTML(http.StatusOK, "users/index.gotmpl", gin.H{ //模板渲染
			"title": "<a href='https://liwenzhou.com'>李文周的博客</a>",
		})

	})

	r.GET("/home", func(c *gin.Context) {
		//http请求
		c.HTML(http.StatusOK, "home.html", nil)

	})

	//启动server
	r.Run(":9090")
}
