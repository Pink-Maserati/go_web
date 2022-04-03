package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func helloTemplate(w http.ResponseWriter, r *http.Request) {
	//解析模板
	files, err := template.ParseFiles("./lesson00/hello.html")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	//渲染模板
	name := "小王子"
	err = files.Execute(w, name)
	if err != nil {
		fmt.Printf("render template failed, err:%v\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/", helloTemplate)
	//获取到当前的路径
	str, _ := os.Getwd()
	fmt.Println(str)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}

}
