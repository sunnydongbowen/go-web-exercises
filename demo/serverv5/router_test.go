package serverv5

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_router_addRoute(t *testing.T) {
	// 这是一个切片，切片存放的是结构体，这种数据结构类似于py里的列表，列表里放的字典。在py里会
	// 到了go里就不会了？？？换了个写法而已啊，常见的Go中的数据结构的定义自己还是要知道的。
	tests := []struct {
		// 输入
		method string
		path   string
	}{
		{ // 静态匹配
			method: http.MethodGet,
			path:   "/",
		},
		{ // 静态匹配
			method: http.MethodGet,
			path:   "//home",
		},
		{
			method: "乱写方法",
			path:   "/",
		},
		{
			method: http.MethodGet,
			path:   "/user",
		},
	}

	wantRouter := &router{
		trees: map[string]*node{
			http.MethodGet: &node{
				path: "/",
			},
			"乱写方法": &node{
				path: "/",
			},
		},
	}

	res := &router{
		trees: map[string]*node{
			http.MethodGet: &node{
				path: "/",
				children: map[string]*node{
					"user": &node{
						path: "user",
					},
					"home": &node{
						path: "home",
					},
				},
			},
		},
	}
	for _, tc := range tests {
		res.addRoute(tc.method, tc.path, nil)
		//t.Run(tc.name)

	}
	// 断言的是整个树
	assert.Equal(t, wantRouter, res)

}
