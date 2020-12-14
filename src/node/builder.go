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

// BuildLeafNode 构造叶子节点
func BuildLeafNode(
	value interface{},
	prev INode,
) INode {
	switch val := value.(type) {
	case string:
		if strings.HasPrefix(val, "$") {
			return NewRef(val)
		} else if strings.HasPrefix(val, ".") {
			return NewCmd(val)
		} else {
			end := NewCmdEnd()
			reg := NewReg(val)
			end.SetPrev(reg)
			reg.SetNext(end)
			return reg
		}
	case map[interface{}]interface{}:
		return NewDict(val)
	case int:
		return NewBack(val)
	}
	log.Fatalf("%v 不能作为叶子节点\n", value)
	panic("node.BuildLeafNode: 致命错误")
}
