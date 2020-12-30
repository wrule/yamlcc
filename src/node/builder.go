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
		} else if strings.HasPrefix(val, ".not") {
			rst = append(rst, NewNot())
		} else if strings.HasPrefix(val, ".other") {
			rst = append(rst, NewOther())
		} else if strings.HasPrefix(val, ".end") {
			rst = append(rst, NewEnd())
		} else {
			rst = append(rst, NewReg(val))
		}
	case int:
		rst = append(rst, NewBack(val))
	case map[interface{}]interface{}:
		for key, value := range val {
			keyNodes := BuildNodes(key)
			for _, keyNode := range keyNodes {
				// 排除定义引用节点展开后的引用节点（next为.end）
				if keyNode.IsNextsEmpty() {
					valueNodes := BuildNodes(value)
					keyNode.Links(valueNodes)
				}
				rst = append(rst, keyNode)
			}
		}
	case []interface{}:
		for _, item := range val {
			rst = append(rst, BuildNodes(item)...)
		}
	default:
		log.Fatalf("%v %T 不能作为节点\n", value, value)
		panic("node.BuildNodes: 致命错误")
	}
	return rst
}

// Compile 编译
func Compile(value interface{}) *Root {
	root := NewRoot()
	nodes := BuildNodes(value)
	root.Links(nodes)
	return root
}
