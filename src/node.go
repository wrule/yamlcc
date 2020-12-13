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
		node.BeginningOf(text)
	} else if num, ok := next.(int); ok {
		if num > 0 {
			me.PrevN(num).BeginningOf(text)
		} else {
			fmt.Println("\t\t\t此路径匹配结束")
		}
	}
}

// TestSuccess 匹配成功提示
func (me *Node) TestSuccess(sub, text string) {
	fmt.Printf("\t\t\t匹配成功: 匹配了 %d 个字符，之后的样子为:%s\n", len(sub), text)
}

// TestFail 匹配失败提示
func (me *Node) TestFail(key string) {
	fmt.Printf("\t\t\t匹配失败: %s\n", key)
}

// BeginningOf 判断文本是否以节点定义为开始
func (me *Node) BeginningOf(text string) (string, string) {
	fmt.Printf(">> 进入新节点，待验证文本:%s\n", text)
	ivdStr, text := me.GetDefRegexp("invalid").StartsWith(text)
	fmt.Printf("\t1. 跳过了 %d 个无效字符，之后的样子为:%s\n", len(ivdStr), text)

	fmt.Printf("\t2. 以下是本节点的分支:\n\t\t")
	for key := range me.childs {
		if !strings.HasPrefix(key, ":") {
			fmt.Printf("%s\t\t", key)
		}
	}
	fmt.Println()

	fmt.Printf("\t3. 开始尝试分支匹配\n")
	for key := range me.childs {
		if strings.HasPrefix(key, "$") {
			defKey := key[1:]
			def := me.GetDef(defKey)
			if re, ok := def.(*RegexpEx); ok {
				fmt.Printf("\t\t%s 分支是正则表达式引用: %v\n", key, re)
				reStr, text := re.StartsWith(text)
				if len(reStr) > 0 {
					me.TestSuccess(reStr, text)
					me.Next(key, text)
				} else {
					me.TestFail(key)
				}
			} else if node, ok := def.(*Node); ok {
				fmt.Printf("\t\t%s 分支是节点引用: %v\n", key, node)
				node.BeginningOf(text)

			} else {
				fmt.Printf("\t\t%s 分支啥也不是\n", key)
			}
		} else if strings.HasPrefix(key, ".") {
			fmt.Printf("\t\t%s 分支是命令\n", key)
			// 这里需要执行命令
		} else {
			fmt.Printf("\t\t%s 分支是字面\n", key)
			if strings.HasPrefix(text, key) {
				text = text[len(key):]
				me.TestSuccess(key, text)
				me.Next(key, text)
			} else {
				me.TestFail(key)
				return "", text
			}
		}
	}
	return "", ""
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
func (me *Node) GetDefRegexp(key string) *RegexpEx {
	if rst, ok := me.GetDef(key).(*RegexpEx); ok {
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
		return &RegexpEx{regexp.MustCompile(value.(string))}
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
