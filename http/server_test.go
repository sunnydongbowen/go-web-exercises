package http

import (
	"fmt"
	"html/template"
	"net/http"
	"testing"
)

// 这个代码是Go 并发编程实战书上的内容

type Data struct {
	Msg string
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	data := Data{"hello http"}
	t.Execute(w, data)
}

func handleText(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		vars := r.URL.Query()
		key, ok := vars["key"]
		if ok {
			msg := "hello get " + key[0]
			w.Write([]byte(msg))
		} else {
			w.Write([]byte("hello world!"))
		}
	}

	if r.Method == "POST" {
		r.ParseForm()
		key := r.Form.Get("name")
		msg := "hello post " + key
		w.Write([]byte(msg))
	}
}

func TestHttp(t *testing.T) {
	http.HandleFunc("/", handleText)
	http.HandleFunc("/index", handleIndex)
	// 一定要写端口号，不然程序一下子就运行完了
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
