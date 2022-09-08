package web

// hander，上课演示的demo,http包过度到web框架
import "net/http"

// 给外部调用的Start方法

func Start() {
	http.ListenAndServe(":8081", &Myhandler{})
}

// 自己定义的handler

type Myhandler struct {
}

// 实现handler接口的方法，它就这一个方法
func (m *Myhandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello world"))
}
