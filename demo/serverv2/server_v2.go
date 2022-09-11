package serverv2

import (
	"net/http"
)

//第一种Server定义

type Server interface {
	http.Handler
	Start(addr string) error
}

type HTTPServer struct {
}

type HTTPSServer struct {
	Server
	//HTTPServer
	CertiFIle string
	KeyFile   string
}

// implement Server interface

func (H *HTTPSServer) Start(addr string) error {
	return http.ListenAndServeTLS(addr, H.CertiFIle, H.KeyFile, H)
}

func (H *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello world"))
}

func (H *HTTPServer) Start(addr string) error {
	// 启动前做点事情
	err := http.ListenAndServe(addr, H)
	// 启动后做点事情
	if err != nil {
		return err
	}
	return err
}

// 对外界来说，这种方式依然是调用这种方法，但是对于框架内部来说，之前的listenandServer换成现在的start方法
func start() {
	var s Server = &HTTPServer{}
	s.Start(":8082")
}
