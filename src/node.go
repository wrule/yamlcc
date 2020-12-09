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

// NewNode 构造函数
func NewNode(srcMap map[interface{}]interface{}) *Node {
	cMap := map[string]interface{}{}
	for _key, value := range srcMap {
		key := _key.(string)
		switch value.(type) {
		case string:
			text := value.(string)
			if strings.HasPrefix(text, "^") {
				cMap[key] = regexp.MustCompile(text)
			} else {
				cMap[key] = value
			}
		case map[interface{}]interface{}:
			cMap[key] = NewNode(value.(map[interface{}]interface{}))
		default:
			cMap[key] = value
		}
	}
	return &Node{
		srcMap: cMap,
	}
}
