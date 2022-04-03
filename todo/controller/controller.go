package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_web/todo/models"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
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

//创建todo
func CreateTodo(c *gin.Context) {
	//创建变量
	var todo models.Todo
	//输入跟模型绑定，即从请求中把数据拿出来
	err := c.ShouldBind(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}
	//数据库保存
	err = models.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

//获取所有的todo列表
func GetAllTodos(c *gin.Context) {
	allTodos, err := models.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, allTodos)
	}
}

func DeleteTodoById(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	err := models.DeleteTodoById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"delete": id,
	})
}

func UpdateTodoById(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}

	//获取这条数据
	todo, err := models.GetTodoById(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return

	}
	if err = models.UpdateTodoById(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, todo)
}
