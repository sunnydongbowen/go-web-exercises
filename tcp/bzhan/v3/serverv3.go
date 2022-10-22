package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 本地端口启动服务
	listener, err := net.Listen("tcp", ":8083")
	if err != nil {
		fmt.Println("start tcp server on 127.0.0.1:8083 failed,err:", err)
		return
	}
	// 等待别人来建立连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed,err:", err)
			return
		}
		go processConn(conn)
	}
}

func processConn(conn net.Conn) {
	// 与客户端进行通信
	var tmp [128]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		// 收消息
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failed,err:", err)
		}
		fmt.Println("收到来自远方的消息: " + string(tmp[:n]))
		// 发消息
		fmt.Print("发送给远方的消息:")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			os.Exit(1)
		}
		conn.Write([]byte([]byte(msg)))
	}
}
