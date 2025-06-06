package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Printf("ParseForm() err:%v", err)
		return
	}
	fmt.Fprintf(w, "Post request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {

	/*
		server针对三个路由：
		1./   返回index.html
		2. /hello 返回 hello func  print hello
		3. / form 返回调用form func 返回form.html
	*/
	fileServer := http.FileServer(http.Dir("./static")) //指定file地址，拿到了一个handle
	//ps:FileServer接收一个http定义好目录路径结构体，然后返回一个注册好的handler
	//这个handler的工作：他有一个serveHTTP方法。
	//这个方法的工作：1.将用户request中的路径映射到本地路径中的对应文件中
	//2.读取文件内容
	//3.使用responseWirter写入响应头以及响应体（文件内容)
	//4.将http响应发送回客户端
	http.Handle("/", fileServer) //针对根目录的请求,重定向到我们映射的位置
	//当请求直接指向目录时返回index.html
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
