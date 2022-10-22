package tcp

import (
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	// 开始监端口
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		t.Fatal(err)
	}
	for {
		//这边开始接受连接，同步处理会阻塞在Accept()
		conn, err := listener.Accept()
		if err != nil {
			t.Fatal(err)
		}
		// 开了一个goroutine去处理,从连接里把数据读出来
		go func() {
			handle(conn)
		}()
	}
}

func handle(conn net.Conn) {
	// 收客户端发来的消息
	reqBs := make([]byte, 8)
	_, err := conn.Read(reqBs[:])
	if err != nil {
		conn.Close()
		return
	}
	// 给客户端返回消息
	_, err = conn.Write([]byte("hello world"))
	if err != nil {
		conn.Close()
		return
	}
}
