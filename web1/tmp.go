// 学习template 模版
package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

//p12 学习template模版
func tmp1() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/action", action)
	http.HandleFunc("/action1", action1)
	http.HandleFunc("/action2", action2)
	http.HandleFunc("/funcMap", functime)

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func process(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("www html/tmp.html")
	err := t.Execute(w, "tmp1 嵌入数据")
	if err != nil {
		return
	}
}

// 学习Action
func action(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("www html/action.html")
	err := t.Execute(w, rand.Intn(10) > 5)
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	rand.Seed(time.Now().Unix())

}

// 循环遍历action
func action1(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("www html/action.html")
	daysOfWeek := []string{"mon", "tue", "wed"}
	err := t.Execute(w, daysOfWeek)
	if err != nil {
		return
	}
}

func action2(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("www html/action.html")
	err := t.Execute(w, "hello")
	if err != nil {
		return
	}
}

// 传入函数
func functime(w http.ResponseWriter, _ *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("www html/tmp.html").Funcs(funcMap)
	t.ParseFiles("www html/tmp.html")
	t.Execute(w, time.Now())
}

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}
func main() {
	tmp1()

}
