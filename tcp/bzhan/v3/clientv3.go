package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//1. 与server端建立链接
	conn, err := net.Dial("tcp", ":8083")
	if err != nil {
		fmt.Println("dial 127.0.0.1:8083 falied,err:", err)
		return
	}
	// 发消息
	var tmp [128]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("发送给远方的消息:")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		// send
		conn.Write([]byte([]byte(msg)))
		// 收消息
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failed,err:", err)

		}
		fmt.Println("收到来自远方的消息:" + string(tmp[:n]))
	}
	conn.Close()
}
