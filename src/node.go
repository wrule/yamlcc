package main

import (
	"regexp"
)

// Node 语法规则节点
type Node struct {
	prev   *Node
	childs map[string]interface{}
}

// SetPrev 设置父节点
func (me *Node) SetPrev(node *Node) {
	me.prev = node
}

// compileValue 编译Value
func compileValue(value interface{}, prev *Node) interface{} {
	switch value.(type) {
	case string:
		return regexp.MustCompile(value.(string))
	case map[interface{}]interface{}:
		node := NewNode(value.(map[interface{}]interface{}))
		node.SetPrev(prev)
		return node
	default:
		return value
	}
}

// compileMap 编译Map
func compileMap(srcMap map[interface{}]interface{}, prev *Node) map[string]interface{} {
	dstMap := map[string]interface{}{}
	for key, value := range srcMap {
		dstMap[key.(string)] = compileValue(value, prev)
	}
	return dstMap
}

// NewNode 构造函数
func NewNode(srcMap map[interface{}]interface{}) *Node {
	node := &Node{}
	node.childs = compileMap(srcMap, node)
	return node
}
