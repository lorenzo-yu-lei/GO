package main

import (
	"net/http"
)

//传入文件，返回页面index.html
func web() {
	err := http.ListenAndServe(":8081", http.FileServer(http.Dir("www html")))
	if err != nil {
		return
	}
}

func main() {
	web()

}
