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
			rst = append(rst, NewDef(":"+valTrimmed))
			rst = append(rst, NewRef("$"+valTrimmed))
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
	}
	if len(rst) < 1 {
		log.Fatalf("%v %T 不能为非叶子节点\n", value, value)
		panic("node.BuildNodes: 致命错误")
	}
	return rst
}

func CompileNode(value map[interface{}]interface{}) INode {

}
