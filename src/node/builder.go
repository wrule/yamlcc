package node

import (
	"log"
	"strings"
)

// BuildNodes 根据传入值创建节点
func BuildNodes(value interface{}) []INode {
	rst := []INode{}
	switch val := value.(type) {
	case string:
		if strings.HasPrefix(val, ":$") {
			valTrimmed := val[2:]
			def := NewDef(":" + valTrimmed)
			ref := NewRef("$" + valTrimmed)
			end := NewEnd()
			ref.Link(end)
			rst = append(rst, def, ref)
		} else if strings.HasPrefix(val, ":") {
			rst = append(rst, NewDef(val))
		} else if strings.HasPrefix(val, "$") {
			rst = append(rst, NewRef(val))
		} else if strings.HasPrefix(val, ".") {
			switch val {
			case ".not":
				rst = append(rst, NewNot())
			case ".other":
				rst = append(rst, NewOther())
			case ".end":
				rst = append(rst, NewEnd())
			}
		} else {
			rst = append(rst, NewReg(val))
		}
	case int:
		rst = append(rst, NewBack(val))
	case map[interface{}]interface{}:
		for key, value := range val {
			keyNodes := BuildNodes(key)
			for _, keyNode := range keyNodes {
				valueNodes := BuildNodes(value)
				keyNode.Links(valueNodes)
				rst = append(rst, keyNode)
			}
		}
	case []interface{}:
		for _, item := range val {
			rst = append(rst, BuildNodes(item)...)
		}
	}
	if len(rst) < 1 {
		log.Fatalf("%v %T 不能为节点\n", value, value)
		panic("node.BuildNodes: 致命错误")
	}
	return rst
}
