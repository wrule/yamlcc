package node

import (
	"log"
	"strings"
)

// BuildNode 构造非叶子节点
func BuildNode(
	value interface{},
	prev INode,
) INode {
	return nil
}

// Link 连接两个节点
func Link(prev, next INode) {
	if prev != nil {
		prev.SetNext(next)
	}
	if next != nil {
		next.SetPrev(prev)
	}
}

// BuildLeafNode 构造叶子节点
func BuildLeafNode(
	value interface{},
	prev INode,
) INode {
	var rst INode = nil
	rstIsEnd := false
	switch val := value.(type) {
	case string:
		if strings.HasPrefix(val, "$") {
			rst = NewRef(val)
		} else if strings.HasPrefix(val, ".") {
			cmd := NewCmd(val)
			if cmd.Cmd() == NodeCmdEnd {
				rstIsEnd = true
				rst = cmd
			}
		} else if strings.HasPrefix(val, ":") {
			// 报错预留
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
	Link(prev, rst)
	if !rstIsEnd {
		end := NewCmdEnd()
		Link(rst, end)
	}
	return rst
}
