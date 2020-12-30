package node

import (
	"log"
	"strings"
)

// BuildCmd 创建命令节点
func BuildCmd(value string) INode {
	switch value {
	case ".not":
		return NewNot()
	case ".other":
		return NewOther()
	case ".end":
		return NewEnd()
	default:
		log.Fatalf("%v %T 不是正确的命令\n", value, value)
		panic("node.BuildNodes: 致命错误")
	}
}

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
			rst = append(rst, BuildCmd(val))
		} else {
			rst = append(rst, NewReg(val))
		}
	case int:
		rst = append(rst, NewBack(val))
	case map[interface{}]interface{}:
		for key, value := range val {
			if _, ok := key.(string); !ok {
				log.Fatalf("%v %T Key只能为字符串类型\n", key, key)
				panic("node.BuildNodes: 致命错误")
			}
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

// Compile 编译语法定义
func Compile(value interface{}) *Root {
	root := NewRoot()
	nodes := BuildNodes(value)
	root.Links(nodes)
	return root
}

// Link 链接语法定义
func Link(root *Root) *Root {
	return nil
}
