package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//监听端口
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 3000,
	})
	//处理监听异常
	if err != nil {
		fmt.Println("listen udp failed,err:", err)
		return
	}
	defer conn.Close()

	// 不需要建立链接，直接收发数据
	var data [1024]byte //定义一个字节型数组，存放收到的信息
	for {
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read from udp failed,err:", err)
			return
		}
		// 这样打直接打的是byte的切片，可以看的更底层一些
		//fmt.Println(data[:])
		fmt.Println(string(data[:n]))

		//转成程大写发给客户端
		//if  {
		//
		//}
		reply := strings.ToUpper(string(data[:n]))
		conn.WriteToUDP([]byte(reply), addr)
	}
}
