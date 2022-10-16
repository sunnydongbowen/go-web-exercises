package serverv4

// 端口监听和服务启动分离
import (
	"fmt"
	"net"
	"net/http"
	"time"
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

// Context 的定义会有很多，先定义核心的
type Context struct {
	Req    *http.Request
	Writer http.ResponseWriter
}

// HandleFunc 看看gin里面怎么定义的，抄过来，不差我一个
type HandleFunc func(*Context)

type Server interface {
	http.Handler
	Start(addr string) error
	// Addroute 注册路的核心抽象，这是核心,Get只是语法糖，只是用户的方便
	Addroute(method, path string, handler HandleFunc)
	//引入注册多个,这个地方注册多个的意思，就是在HandleFunc里写多个函数，就是一个url，它其实是调用了多个函数来处理这个逻辑
	//Addroutes(method, path string, handlers... HandleFunc)
}

func (H *HTTPServer) Addroute(method, path string, handler HandleFunc) {

}

// Get 下面这个 其实只是为了方便我们使用，上面三个参数也是一样的,为了使用者方便封装了一下，其实和上面的Addroute是一样的
func (H *HTTPServer) Get(path string, handler HandleFunc) {
	H.Addroute(http.MethodGet, path, handler)
}

func (H *HTTPServer) POST(path string, handler HandleFunc) {
	H.Addroute(http.MethodPost, path, handler)
}

func (H *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// writer.Write([]byte("hello world"))
	ctx := &Context{
		Req:    request,
		Writer: writer,
	}
	// 接下来就是查找路由，执行方法调用，业务逻辑
	H.serve(ctx)
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

func (H *HTTPServer) serve(ctx *Context) {

}

func start() {
	var s Server = &HTTPServer{}
	//var s1 Server =&HTTPSServer{
	//	&HTTPServer{},
	//	"xxx",
	//	"aaa",
	//}
	// 为了演示说，用户可以自己将多个 handleFunc 聚合在一起，已解决顺序问题
	// 用户自己先调用 h1，再调用 h2，自己手动调用，其它框架是允许用户注册多个的，但是我们不允许，所以我解释用户可以自己封装
	var h1 HandleFunc = func(context *Context) {
		fmt.Println("步骤1")
		time.Sleep(time.Second)
	}

	var h2 HandleFunc = func(context *Context) {
		fmt.Println("步骤2")
		time.Sleep(time.Second)
	}

	s.Addroute(http.MethodPost, "/user", func(context *Context) {
		// 循环调用多个
		h1(context)
		h2(context)
	})

	s.Start(":8082")
}
