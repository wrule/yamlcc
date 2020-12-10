package main

import (
	"regexp"
	"strings"
)

// Node s
type Node struct {
	srcMap map[string]interface{}
}

// Keys s
func (me *Node) Keys() []string {
	rst := []string{}
	for key := range me.srcMap {
		rst = append(rst, key)
	}
	return rst
}

// Values s
func (me *Node) Values() []interface{} {
	rst := []interface{}{}
	for _, value := range me.srcMap {
		rst = append(rst, value)
	}
	return rst
}

// Child 获取子节点
func (me *Node) Child(key string) interface{} {
	if value, found := me.srcMap[key]; found {
		return value
	}
	panic("Node Child: 无法获取子节点")
}

// ChildNode 获取子节点Node
func (me *Node) ChildNode(key string) Node {
	if node, ok := me.Child(key).(Node); ok {
		return node
	}
	panic("Node ChildNode: 不是一个节点")
}

// ChildString 获取子节点字符串
func (me *Node) ChildString(key string) string {
	if str, ok := me.Child(key).(string); ok {
		return str
	}
	panic("Node ChildString: 不是一个字符串")
}

// ChildRegexp 获取子正则表达式
func (me *Node) ChildRegexp(key string) *regexp.Regexp {
	if re, ok := me.Child(key).(*regexp.Regexp); ok {
		return re
	}
	panic("Node ChildString: 不是一个正则表达式")
}

// compileValue 编译Value
func compileValue(value interface{}) interface{} {
	switch value.(type) {
	case string:
		text := value.(string)
		if strings.HasPrefix(text, "^") {
			return regexp.MustCompile(text)
		}
		return value
	case map[interface{}]interface{}:
		childMap := value.(map[interface{}]interface{})
		return NewNode(childMap)
	default:
		return value
	}
}

// compileMap 编译Map
func compileMap(srcMap map[interface{}]interface{}) map[string]interface{} {
	dstMap := map[string]interface{}{}
	for key, value := range srcMap {
		dstMap[key.(string)] = compileValue(value)
	}
	return dstMap
}

// NewNode 构造函数
func NewNode(srcMap map[interface{}]interface{}) *Node {
	return &Node{
		srcMap: compileMap(srcMap),
	}
}
