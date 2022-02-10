package main

import "fmt"

type node struct {
	// 节点名字
	name string
	// 子节点
	children map[string]*node
	// 子节点是否是动态的
	childDyn string
	// 参数名称
	params string
	// 是否是最终节点
	end bool
}

func newNode(name string) *node {
	return &node{
		name:     name,
		children: make(map[string]*node),
		params:   "",
		end:      false,
	}
}

// 根据name获取子节点
// 如果有就返回
// 如果没有就添加到子节点map里面，并返回
func (n *node) getChilds(name string) *node {
	if name[0] == '*' || name[0] == ':' {
		if n.childDyn != "" {
			fmt.Println("匹配到动态参数,n.name=", n.name, "  name=>", name)
			return n.children[n.childDyn]
		} else {
			n.childDyn = name
		}

	}

	node, ok := n.children[name]
	if ok {

		return node
	} else {
		node = newNode(name)
		n.children[name] = node
		return node
	}

}

func (n *node) searchNodes(name string) (*node, bool) {
	if n.children == nil {
		return nil, false
	}

	if n.childDyn != "" {
		var temp *node
		temp = n.children[n.childDyn]
		temp.params = name
		return temp, true

	}

	node, ok := n.children[name]
	if ok {

		return node, ok
	} else {
		return nil, false
	}
}
