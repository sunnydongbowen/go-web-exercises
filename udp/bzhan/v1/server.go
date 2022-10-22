package main

import (
	"fmt"
	"net"
)

func main() {
	// 监听
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 3000,
	})
	if err != nil {
		fmt.Println("listen failed,err:", err)
		return
	}
	defer listener.Close()

	for {
		//接收数据
		var data [1024]byte
		n, addr, err := listener.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed,err:", err)
			continue
		}
		fmt.Printf("data:%v addr:%v count：%v\n", string(data[:n]), addr, n)

		// 发送数据,这里把接收的又发给客户端了！
		_, err = listener.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Println("write to udp failed,err:", err)
			continue
		}
	}
}
