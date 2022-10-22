package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 连接
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
	})
	if err != nil {
		fmt.Println("dial failed,err:", err)
		return
	}
	defer socket.Close()

	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		// 发数据
		fmt.Print("请输入内容:")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		socket.Write([]byte(msg))

		n, _, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("rec reply failed,err:", err)
			return
		}
		fmt.Println("收到回复的信息: ", string(reply[:n]))
	}
}
