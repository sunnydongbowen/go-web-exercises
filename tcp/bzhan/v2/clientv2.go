package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 1. 与server端建立连接
	conn, err := net.Dial("tcp", ":8082")
	if err != nil {
		fmt.Println("dial 127.0.0.1:8082 falied,err:", err)
		return
	}
	// 发送数据
	var msg string
	if len(os.Args) <= 2 {
		msg = "hello bowen"
	} else {
		msg = os.Args[1]
	}
	conn.Write([]byte([]byte(msg)))
	conn.Close()
}
