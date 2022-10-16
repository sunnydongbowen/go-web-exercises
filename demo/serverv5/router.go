package serverv5

import (
	"strings"
)

// 代表路由
type router struct {
	// gin trees []*tree
	// get。post等都需要有一颗树
	trees map[string]*node
}

// 树
//type tree struct {
//	root *node
//}

// 命中了这里，就去执行方法

type node struct {
	//  /a/b/c 中的b这一段
	path     string
	handler  HandleFunc
	children map[string]*node
	// children map[string]*node
}

func (r *router) addRoute(method string, path string, handleFunc HandleFunc) {
	tree, ok := r.trees[method]
	if !ok {
		// 根节点
		tree = &node{path: "/"}
		r.trees[method] = tree
	}

	if path == "/" {
		return
	}

	// 把前后的/都去掉
	path = strings.Trim(path, "/")

	// 支持/user
	segs := strings.Split(path, "/")
	//println(segs)

	if tree.children == nil {
		tree.children = make(map[string]*node)
	}
	tree.children[segs[0]] = &node{path: segs[0], handler: handleFunc}

}
