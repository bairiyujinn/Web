package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//注册请求路径的handler的函数
	//设置两个路由，分别绑定indexHandler,helloHandler
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)

	//log.Fatal函数是记录致命错误并终止程序的执行
	//http.ListenAndServe启动端口为9999的http服务
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// http.ResponseWriter用于在 HTTP 服务器端向客户端发送响应
// http.Request用于封装客户端发来的 HTTP 请求相关的所有信息,包括请求方法和URL
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// 通过 r.Header 来获取请求携带的各种头部信息，比如获取 User-Agent 字段来了解客户端的类型等。
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}
}