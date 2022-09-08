package serverv1

import "net/http"

// 自己定义的server

type Server interface {
	http.Handler
}

// 自己定义的结构体类型，实现http server	,

type HTTPServer struct {
}

func (H *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello world"))
}

// 别人要用，直接调用我们提供的start方法就好

func Start() {
	var s Server = &HTTPServer{}
	// 可以启动普通服务器
	http.ListenAndServe(":8081", s)
	// 无缝连接
	http.ListenAndServeTLS("4000", "xxx", "aaa", s)

}
