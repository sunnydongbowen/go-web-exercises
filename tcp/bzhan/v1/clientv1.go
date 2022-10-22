package main

import (
	"fmt"
	"net"
)

func main() {
	//1. 与server端建立链接
	conn, err := net.Dial("tcp", ":8082")
	if err != nil {
		fmt.Println("dial 127.0.0.1:8082 falied,err:", err)
		return
	}
	conn.Write([]byte("hello bowen"))
	conn.Close()
}
