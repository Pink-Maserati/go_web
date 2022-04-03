package main

import (
	"fmt"
	html "html/template"
	"net/http"
	text "text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := html.New("index.gotmpl").Delims("{[", "]}").ParseFiles("./index.gotmpl")

	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	name := "小王子"
	t.Execute(w, name)

}

func xss(w http.ResponseWriter, r *http.Request) {
	t, err := text.ParseFiles("./xss.gotmpl")

	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}

	str := "<script>alert('嘿嘿嘿')</script>"
	t.Execute(w, str)

}

func xss2(w http.ResponseWriter, r *http.Request) {

	//解析模板前定义一个自定义的函数safe
	t, err := html.New("xss.gotmpl").Funcs(html.FuncMap{
		"safe": func(str string) html.HTML {
			return html.HTML(str)

		},
	}).ParseFiles("./xss.gotmpl")

	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}

	str1 := "<script>alert('嘿嘿嘿')</script>"
	str2 := "<a href='http://liwenzhou.com'>liwenzhou></a>"
	t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})

}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	http.HandleFunc("/xss2", xss2)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
