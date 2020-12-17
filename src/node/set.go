package node

import "fmt"

// Set 集合节点
type Set struct {
	nodeSet []INode
	Com
}

// BeginningOf 匹配
func (me *Set) BeginningOf(text string) (string, string, bool) {
	return "", text, true
}

func getNodeSet(srcSet []interface{}, prev INode) []INode {
	rst := []INode{}
	for _, item := range srcSet {
		node := BuildLeafNode(item)
		node.SetPrev(prev)
		rst = append(rst, node)
	}
	return rst
}

// NewSet 构造函数
func NewSet(srcSet []interface{}) *Set {
	rst := &Set{}
	rst.nodeSet = getNodeSet(srcSet, rst)
	rst.Com = Com{me: rst, srcValue: srcSet}
	return rst
}

// Print 打印节点信息
func (me *Set) Print() {
	fmt.Printf("集合节点: %v\n", me.SrcValue())
}
