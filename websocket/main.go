package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"net/http"
)

// websocket server
func main() {
	fmt.Println("websocket at localhost:8080/echo")
	// 绑定handler
	http.Handle("/echo", websocket.Handler(Echo))
	http.HandleFunc("/", handleIndex)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

type Data struct {
	Msg string
}

// 返回静态html
func handleIndex(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("websocket/index.html")
	//data := Data{msg}
	//t.Execute(writer, data)
	t.Execute(writer, nil)
}

func Echo(w *websocket.Conn) {
	var err error
	for {
		var recMsg string
		err = websocket.Message.Receive(w, &recMsg)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("客户端: " + recMsg)
		msg := ""
		if recMsg == "猜猜年龄" {
			msg = "服务器：18岁"
		} else if recMsg == "你好" {
			msg = "服务器:你好，请问有什么可以帮到你？"
		} else {
			msg = "服务器" + recMsg
		}
		fmt.Println(msg)
		//给客户端发消息
		if err = websocket.Message.Send(w, msg); err != nil {
			fmt.Println(err)
			break
		}
	}
}
