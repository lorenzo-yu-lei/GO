package main

import (
	"fmt"
	"log"
	"net/http"
)

//创建一个helloworld 页面
func web1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("/根目录"))
		if err != nil {
			return
		}
	})
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}

//创建一个自己的handle
type myHandler struct{}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("加油 雷宇 GO GO GO"))
	if err != nil {
		return
	}
}

type aboutHandler struct{}

func (m *aboutHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("about!"))
	if err != nil {
		return
	}
}

func good(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("加油，你很棒"))
	if err != nil {
		return
	}
}

func handlefunc(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("http.Handlefunc() 是类型也是函数"))
	if err != nil {
		return
	}
}

// 处理handle请求
func web2() {
	m := myHandler{}
	a := aboutHandler{}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	http.Handle("/hello", &m)
	http.Handle("/about", &a)

	//http.HandleFunc 方法
	http.HandleFunc("/good", good)
	http.Handle("/Handlefunc", http.HandlerFunc(handlefunc))
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func query() {
	http.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL
		query := url.Query()

		id := query["id"]
		log.Println(id)

		name := query.Get("name")
		log.Println(name)
		fmt.Print("1123341")
	})
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}

func main() {
	//web1()
	web2()
	query()
}
