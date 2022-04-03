package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//定义一个函数kua 要么只有一个返回值，要么有两个返回值，第二个返回值必须是error类型
	kua := func(name string) (string, error) {
		return name + "好帅", nil
	}
	//创建一个名字是f1的模板函数，名字一定要与模板的名字能对应上
	t, err := template.New("f1.tmpl").
		Funcs(template.FuncMap{
			"kua": kua, //嘎岁模板引擎，现在多了一个自定义函数kua
		}).
		ParseFiles("./f1.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	name := "小王子"
	t.Execute(w, name)
}

func f2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}

	t.Execute(w, "小王子")

}
func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmpl", f2)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
