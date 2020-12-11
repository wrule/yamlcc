package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Node 语法规则节点
type Node struct {
	name   string
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

// Next s
func (me *Node) Next(key, text string) {
	next := me.childs[key]
	if node, ok := next.(*Node); ok {
		node.Test(text)
	} else if num, ok := next.(int); ok {
		if num > 0 {
			me.PrevN(num).Test(text)
		} else {
			fmt.Println("\t\t\t此路径匹配结束")
		}
	}
}

func (me *Node) RegexpFind(re *regexp.Regexp, text string) (string, string) {
	indexs := re.FindStringIndex(text)
	if len(indexs) > 1 && indexs[0] == 0 {
		dst := text[indexs[0]:indexs[1]]
		return dst, text[len(dst):]
	}
	return "", text
}

// Test 验证文本
func (me *Node) Test(text string) {
	fmt.Printf(">> 进入新节点\n")
	fmt.Printf(">> 待验证文本:%s\n", text)

	ivdStr, text := me.RegexpFind(me.GetDefRegexp("invalid"), text)
	fmt.Printf("\t1. 跳过了 %d 个无效字符，之后的样子为:%s\n", len(ivdStr), text)

	fmt.Printf("\t2. 以下是本节点的分支:\n")
	for key := range me.childs {
		fmt.Printf("\t\t%s\t", key)
		if strings.HasPrefix(key, "$") {
			fmt.Println("引用")
		} else if strings.HasPrefix(key, ".") {
			fmt.Println("命令")
		} else {
			fmt.Println("字面")
		}
	}

	fmt.Printf("\t3. 开始尝试分支匹配\n")
	for key := range me.childs {
		fmt.Printf("\t\t尝试匹配分支:%s\n", key)
		if strings.HasPrefix(key, "$") {
			fmt.Printf("\t\t\t%s 分支是引用\n", key)

			defKey := key[1:]
			def := me.GetDef(defKey)

			if re, ok := def.(*regexp.Regexp); ok {
				fmt.Printf("\t\t\t引用是正则表达式: %v\n", re)
				reStr, text := me.RegexpFind(re, text)
				if len(reStr) > 0 {
					fmt.Println("\t\t\t匹配成功")
					fmt.Printf("\t\t\t匹配到 %s，共 %d 个字符，之后的样子为:%s\n", reStr, len(reStr), text)
					me.Next(key, text)
				} else {
					fmt.Println("\t\t\t匹配不成功")
				}

			} else if node, ok := def.(*Node); ok {
				fmt.Printf("\t\t\t引用是节点: %v\n", node)
				node.Test(text)
			} else {
				fmt.Printf("\t\t\t引用啥也不是\n")
			}

		} else if strings.HasPrefix(key, ".") {
			fmt.Printf("\t\t\t%s 分支是命令\n", key)

		} else {
			fmt.Printf("\t\t\t%s 分支是字面\n", key)
			if strings.HasPrefix(text, key) {
				fmt.Println("\t\t\t匹配成功")
				text = text[len(key):]
				fmt.Printf("\t\t\t匹配了 %d 个字符，之后的样子为:%s\n", len(key), text)

				me.Next(key, text)

			} else {
				fmt.Println("\t\t\t匹配不成功")
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
