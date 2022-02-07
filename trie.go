package main

import (
	"fmt"
	"strings"
)

type Trie struct {

	// 子节点
	root *node
}

func NewTrie() *Trie {
	return &Trie{

		root: newNode("/"),
	}
}

func (t *Trie) Insert(patterns string) {

	path := strings.Split(patterns, "/")
	var node = t.root
	for k, v := range path {
		node = node.getChilds(v)
		fmt.Println(k, " node=", node)
	}
	node.end = true
}

func getChilds(node *node, name string) *node {
	if node == nil {
		node = newNode(name)
		return node
	}
	n, ok := node.children[name]
	if !ok {
		n = newNode(name)
		node.children[name] = n

	}
	return n
}

func (t *Trie) Search(patterns string) (*node, bool) {
	if patterns == "" {
		return nil, false
	}
	path := strings.Split(patterns, "/")
	if len(path) == 0 {
		return nil, false
	}
	var node = t.root
	// api/v1/java/doc
	for _, v := range path {
		var ok bool
		node, ok = node.searchNodes(v)
		if !ok {
			return nil, false
		}
	}
	// 判断是否是注册的url
	if node.end == false {
		return nil, false
	}
	return node, true
}
