package main

//query string
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//GET请求 URL ?后面的是querystring参数
	//key=value格式，多个key-value用&连接
	//eq: web?query=123sb&age=18
	r.GET("/web", func(c *gin.Context) {
		//或哦去浏览器那边发请求携带的query string参数
		//方式1：通过Query获取请求中携带的querystring参数
		name := c.Query("query")
		age := c.Query("age")
		//方式2：取不到就用默认值
		//name:=c.DefaultQuery("query","somebody")
		//方式3：取到返回（值，true) 取不到返回（""，false)
		//name,ok:=c.GetQuery("query")
		//if !ok{
		//	name="somebody ok"
		//}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})

	})
	r.Run(":9090")

}
