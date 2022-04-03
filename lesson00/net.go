package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//file, _ := ioutil.ReadFile("./hello.txt")
	file, _ := ioutil.ReadFile("./lesson00/hello.html")
	_, _ = fmt.Fprintln(w, string(file))

}

func main() {

	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}

}
