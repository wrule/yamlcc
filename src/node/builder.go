package node

import (
	"log"
	"strings"
)

// Link 连接两个节点
func Link(prev, next INode) {
	if prev != nil {
		prev.SetNext(next)
	}
	if next != nil {
		next.SetPrev(prev)
	}
}

// BuildNode 构造非叶子节点
// 定义，引用，非结束命令，正则可以作为非叶子节点
func BuildNode(
	value interface{},
) INode {
	var rst INode = nil
	switch val := value.(type) {
	case string:
		if strings.HasPrefix(val, ":") {
			rst = NewDef(val)
		} else if strings.HasPrefix(val, "$") {
			rst = NewRef(val)
		} else if strings.HasPrefix(val, ".") {
			cmd := NewCmd(val)
			if cmd.Cmd() != NodeCmdEnd {
				rst = cmd
			}
		} else {
			rst = NewReg(val)
		}
	}
	if rst == nil {
		log.Fatalf("%v 不能为非叶子节点\n", value)
		panic("node.BuildNode: 致命错误")
	}
	return rst
}

// BuildLeafNode 构造叶子节点
// 正则表达式，引用，字典，回跳，结束命令可以作为叶子节点
func BuildLeafNode(
	value interface{},
) INode {
	var rst INode = nil
	rstIsEnd := false
	switch val := value.(type) {
	case string:
		if strings.HasPrefix(val, ":") {
			// 报错预留
		} else if strings.HasPrefix(val, "$") {
			rst = NewRef(val)
		} else if strings.HasPrefix(val, ".") {
			cmd := NewCmd(val)
			if cmd.Cmd() == NodeCmdEnd {
				rstIsEnd = true
				rst = cmd
			}
		} else {
			rst = NewReg(val)
		}
	case map[interface{}]interface{}:
		rst = NewDict(val)
	case int:
		rst = NewBack(val)
	}
	if rst == nil {
		log.Fatalf("%v 不能为叶子节点\n", value)
		panic("node.BuildLeafNode: 致命错误")
	}
	if !rstIsEnd {
		end := NewCmdEnd()
		Link(rst, end)
	}
	return rst
}
