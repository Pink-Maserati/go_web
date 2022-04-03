package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func hello(w http.ResponseWriter, r *http.Request) {
	//解析模板
	files, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	//渲染模板
	u := User{
		Name:   "小王子",
		Gender: "男",
		Age:    20,
	}
	m := map[string]interface{}{
		"name":   "小仙女",
		"gender": "女",
		"age":    22,
	}
	hobbies := []string{"钢琴", "英语"}
	files.Execute(w, map[string]interface{}{
		"u":       u,
		"m":       m,
		"hobbies": hobbies,
	})

}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
