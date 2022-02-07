package main

type node struct {
	// 节点名字
	name string
	// 子节点
	children map[string]*node
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

func (n *node) getChilds(name string) *node {

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
	node, ok := n.children[name]
	if ok {

		return node, ok
	} else {
		return nil, false
	}
}
