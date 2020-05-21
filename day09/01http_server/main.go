package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// http_server

func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./xx.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v\n", err)))
	}
	w.Write(b) // 类型强制转换---类型名+（）  []byte是类型，[]byte(str),把字符串强制转换为字节类型的变量
}

func f2(w http.ResponseWriter, r *http.Request) {
	// 对于GET请求，参数都放在URL上，（query param），请求体中是没有数据的
	queryParam := r.URL.Query() // 自动帮我们识别URL中的query param
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/xxx/", f1)
	http.HandleFunc("/x/xx/", f2)
	http.ListenAndServe("127.0.0.1:9000", nil)
}
