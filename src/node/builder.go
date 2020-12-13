package node

import (
	"fmt"
	"strings"
)

// BuildNode 构造节点的方法
func BuildNode(
	value interface{},
	prev INode,
	next INode,
	key bool,
) INode {
	switch val := value.(type) {
	case string:
		fmt.Println("字符串")
		if strings.HasPrefix(val, ":") {
			return NewDef(val, prev, next)
		} else if strings.HasPrefix(val, "$") {
			return NewRef(val, prev, next)
		} else if strings.HasPrefix(val, ".") {
			return NewCmd(val, prev, next)
		} else {
			return NewLit(val, prev, next)
		}
	case map[interface{}]interface{}:
		fmt.Println("字典")
		return NewDict(val, prev)
	case int:
		if val > 0 {
			return NewBack(val, prev, nil)
		}
		return NewEnd(prev)
	default:
		panic("类型错误\n")
	}
}
