package node

import (
	"log"
	"strings"
)

// BuildNodes 构造非叶子节点
// 定义引用，定义，引用，命令（.other，.not），正则，可以作为非叶子节点
// 定义引用节点会产生两个非叶子节点，一个为定义一个为引用（后跟.end节点）
func BuildNodes(
	value interface{},
) []INode {
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
			case ".other":
				rst = append(rst, NewOther())
			case ".not":
				rst = append(rst, NewNot())
			}
		} else {
			rst = append(rst, NewReg(val))
		}
	}
	if len(rst) < 1 {
		log.Fatalf("%v %T 不能为非叶子节点\n", value, value)
		panic("node.BuildNodes: 致命错误")
	}
	return rst
}

// BuildLeafNode 构造叶子节点
// 引用，命令（.end），正则，字典，集合，回跳，可以作为叶子节点
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
			switch val {
			case ".end":
				rstIsEnd = true
				rst = NewEnd()
			}
		} else {
			rst = NewReg(val)
		}
	case map[interface{}]interface{}:
		rst = NewDict(val)
	case []interface{}:
		rst = NewSet(val)
	case int:
		rst = NewBack(val)
	}
	if rst == nil {
		log.Fatalf("%v %T 不能为叶子节点\n", value, value)
		panic("node.BuildLeafNode: 致命错误")
	}
	if !rstIsEnd {
		end := NewEnd()
		rst.Link(end)
	}
	return rst
}
