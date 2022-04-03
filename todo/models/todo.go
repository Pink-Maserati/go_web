package models

import (
	"go_web/todo/dao"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// CreateATodo 创建todo
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Debug().Create(&todo).Error
	return
}

//获取所有的todos
func GetAllTodos() (todos []*Todo, err error) {
	if err = dao.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return
}

func DeleteTodoById(id string) (err error) {
	err = dao.DB.Delete(&Todo{}, id).Error
	return
}

func GetTodoById(id string) (todo *Todo, err error) {
	//todo = new(Todo)
	if err = dao.DB.Debug().Where("id=?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return
}
func UpdateTodoById(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}
