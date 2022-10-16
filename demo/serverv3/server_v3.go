package serverv3

// 端口监听和服务启动
import (
	"net"
	"net/http"
)

// alt+insert，implement the interface

type HTTPServer struct {
}

// 装饰器模式

type HTTPSServer struct {
	Server
	//HTTPServer
	CertFile string
	KeyFile  string
}
type Server interface {
	http.Handler
	Start(addr string) error
}

func (H *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello world"))
}

func (H *HTTPServer) Start(addr string) error {
	// 端口启动前
	//var listener net.Listener
	// 这里只是启动了端口，客户发发请求的话是得不到响应的
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	// 端口启动后回调 是你启动了端口之后，你就可以做点别的事情了,
	// 注册本服务器到我的管理平台，比如说通知你的负载均衡，通知你的 admin 后台
	// 比如注册到etcd，然后打开管理界面，就能看到这个实例 10.0.0。1:8081
	println("成功监听端口")
	return http.Serve(listener, H) // 只有到了这，客户端发请求才会得到响应
	// 这种方式是直接阻塞的，没办法做其它事情了，启动后就阻塞在这里了，不阻塞的话程序就退出来了，这显然不符合服务程序的设计方式。
	// return http.ListenAndServe(addr,H)
}

func start() {
	var s Server = &HTTPServer{}
	s.Start(":8082")
}
