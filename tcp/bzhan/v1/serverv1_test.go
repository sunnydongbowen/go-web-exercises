package main

import (
	"fmt"
	"net"
	"testing"
)

func TestServer1(t *testing.T) {
	// 1. 本地端口启动服务
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		fmt.Println("start tcp server on 127.0.0.1:8082 failed,err:", err)
		return
	}
	//2. 等待别人来跟我建立链接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("accept failed,err:", err)
		return
	}
	//3. 与客户端通信
	var tmp [128]byte
	n, err := conn.Read(tmp[:])
	if err != nil {
		fmt.Println("read from conn failed,err:", err)
	}
	fmt.Println(string(tmp[:n]))
}
