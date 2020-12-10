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
	keys   []string
}

// Get 访问某定义（会向上查找）
func (me *Node) Get(key string) interface{} {
	curNode := me
	for curNode != nil {
		if rst, ok := curNode.srcMap[":"+key]; ok {
			return rst
		}
		curNode = curNode.prev
	}
	panic("Node Get: 定义没有访问到")
}

// GetNode 获取子节点定义
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
	invalid := me.GetRegexp("invalid").FindString(text)
	fmt.Printf("跳过%d个空白符\n", len(invalid))
	text = text[len(invalid):]

	fmt.Println(text)

	for _, key := range me.keys {
		fmt.Printf("%s\t", key)
		if strings.HasPrefix(key, "$") {
			fmt.Printf("尝试匹配引用: %s\n", key)
			yy := me.Get(key[1:])
			if re, ok := yy.(*regexp.Regexp); ok {
				fmt.Printf("引用是一个正则表达式: %v\n", re)
				kkk := re.FindString(text)
				if len(kkk) > 0 {
					fmt.Printf("匹配成功，为: %s\n", kkk)
					next := me.srcMap[key]
					// fmt.Println(next)
					text = text[len(kkk):]
					if node, ok := next.(*Node); ok {
						node.Run(text)
					} else {
						return
					}
				} else {
					fmt.Printf("匹配失败\n")
				}
			} else if node, ok := yy.(*Node); ok {
				fmt.Printf("引用是一个组合定义: %v\n", node)
			}
		} else {
			fmt.Printf("尝试匹配字符串: %s\n", key)
			if strings.HasPrefix(text, key) {
				fmt.Printf("匹配成功\n")
				next := me.srcMap[key]
				// fmt.Println(next)
				text = text[len(key):]
				if node, ok := next.(*Node); ok {
					node.Run(text)
				} else {
					return
				}
			} else {
				fmt.Printf("匹配失败\n")
			}
		}
	}
}

// SetPrev s
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
	node.srcMap = compileMap(srcMap, node)
	for key := range node.srcMap {
		node.keys = append(node.keys, key)
	}
	return node
}
