package main

//文件上传
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html", "./index2.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/multiIndex", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index2.html", nil)
	})

	//单个文件的上传
	r.POST("/upload", func(c *gin.Context) {
		//从请求中读取文件
		f, err := c.FormFile("file1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			//将读取到的文件保存到本地（服务端本地）
			//dst:=fmt.Sprintf("./%s",f.Filename)
			dst := path.Join("./", f.Filename)
			c.SaveUploadedFile(f, dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	r.POST("/multiUpload", func(c *gin.Context) {

		//多个文件上传
		form, _ := c.MultipartForm()
		files := form.File["file1"]
		fmt.Printf("value:%#v\n", files)
		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("./%s%d", file.Filename, index)
			c.SaveUploadedFile(file, dst)

		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})

	r.Run(":9090")

}
