package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Node s
type Node struct {
	srcMap map[string]interface{}
	prev   *Node
}

// Keys 子字段列表
func (me *Node) Keys() []string {

}

// ChildStrings 子字符串列表
func (me *Node) ChildStrings() []string {

}

// ChildRegexps 子正则表达式列表
func (me *Node) ChildRegexps() []*regexp.Regexp {

}

// ChildNodes 子节点列表
func (me *Node) ChildNodes() []*Node {

}

// Get 访问某字段（会向上查找）
func (me *Node) Get(key string) interface{} {
	curNode := me
	for curNode != nil {
		if rst, ok := curNode.srcMap[key]; ok {
			return rst
		}
		curNode = curNode.prev
	}
	panic("Node Get: 字段没有访问到")
}

// GetNode 获取子节点Node
func (me *Node) GetNode(key string) *Node {
	if node, ok := me.Get(key).(*Node); ok {
		return node
	}
	panic("Node GetNode: 不是一个节点")
}

// GetString 获取子节点字符串
func (me *Node) GetString(key string) string {
	if str, ok := me.Get(key).(string); ok {
		return str
	}
	panic("Node GetString: 不是一个字符串")
}

// GetRegexp 获取子正则表达式
func (me *Node) GetRegexp(key string) *regexp.Regexp {
	if re, ok := me.Get(key).(*regexp.Regexp); ok {
		return re
	}
	panic("Node GetRegexp: 不是一个正则表达式")
}

// Run 运行
func (me *Node) Run(text string) {
	istr := me.GetRegexp("$invalid").FindString(text)
	fmt.Println(text, len(istr))
}

// SetPrev s
func (me *Node) SetPrev(node *Node) {
	me.prev = node
}

// compileValue 编译Value
func compileValue(value interface{}, prev *Node) interface{} {
	switch value.(type) {
	case string:
		text := value.(string)
		if strings.HasPrefix(text, "^") {
			return regexp.MustCompile(text)
		}
		return value
	case map[interface{}]interface{}:
		childMap := value.(map[interface{}]interface{})
		node := NewNode(childMap)
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
	node.srcMap = compileMap(srcMap, node)
	return node
}
