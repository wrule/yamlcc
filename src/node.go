package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Node 语法规则节点
type Node struct {
	prev   *Node
	childs map[string]interface{}
}

// PrevN 获取多级父节点
func (me *Node) PrevN(n int) *Node {
	curNode := me
	for i := 0; curNode != nil && i < n; i++ {
		curNode = curNode.prev
	}
	return curNode
}

// Test 验证文本
func (me *Node) Test(text string) {
	fmt.Println("->>进入新节点")
	for key := range me.childs {
		fmt.Printf("%s\t", key)
	}
	fmt.Println()
	fmt.Println(text)
	fmt.Println("--------------------")
	fstr := me.GetDefRegexp("invalid").FindString(text)
	fmt.Printf("无效字符 %d 个\n", len(fstr))
	text = text[len(fstr):]
	fmt.Println(text)
	for key := range me.childs {
		fmt.Printf("分支: %s ", key)
		if strings.HasPrefix(key, "$") {
			defKey := key[1:]
			def := me.GetDef(defKey)
			if re, ok := def.(*regexp.Regexp); ok {
				fmt.Println("是正则定义", re)
				reStr := re.FindString(text)
				if reStr != "" {
					fmt.Printf("正则匹配成功: %s %d个字符\n", reStr, len(reStr))
					text = text[len(reStr):]
					fmt.Println(text)
					next := me.childs[key]
					if node, ok := next.(*Node); ok {
						node.Test(text)
					} else {
						panic("结束")
					}
				} else {
					fmt.Println("正则匹配失败")
				}
			} else if node, ok := def.(*Node); ok {
				fmt.Println("是节点定义", node)
				node.Test(text)
			}
		} else {
			fmt.Println("是字面量")
			if strings.HasPrefix(text, key) {
				fmt.Println("匹配成功")
				text = text[len(key):]
				fmt.Println(text)
				next := me.childs[key]
				if node, ok := next.(*Node); ok {
					node.Test(text)
				} else {
					panic("结束")
				}
			} else {
				fmt.Println("不匹配")
			}
		}
	}
}

// GetDef 获取定义
func (me *Node) GetDef(key string) interface{} {
	curNode := me
	for curNode != nil {
		if rst, ok := curNode.childs[":"+key]; ok {
			return rst
		}
		curNode = curNode.prev
	}
	panic("Node GetDef: 找不到定义 " + key)
}

// GetDefRegexp 获取正则表达式定义
func (me *Node) GetDefRegexp(key string) *regexp.Regexp {
	if rst, ok := me.GetDef(key).(*regexp.Regexp); ok {
		return rst
	}
	panic("Node GetDefRegexp: 不是正则表达式定义 " + key)
}

// GetDefNode 获取语法节点定义
func (me *Node) GetDefNode(key string) *Node {
	if rst, ok := me.GetDef(key).(*Node); ok {
		return rst
	}
	panic("Node GetDefNode: 不是语法节点定义 " + key)
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
	node.SetPrev(nil)
	return node
}
