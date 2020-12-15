package node

import "fmt"

// Com 节点共用部分
type Com struct {
	srcValue interface{}
	prev     INode
	next     INode
}

// SrcValue 原始值
func (me *Com) SrcValue() interface{} {
	return me.srcValue
}

// Prev 获取上一个节点
func (me *Com) Prev() INode {
	return me.prev
}

// PrevN 获取上n个节点（n大于等于1）
func (me *Com) PrevN(n int) INode {
	curNode := me.Prev()
	for i := 1; i < n && curNode != nil; i++ {
		curNode = curNode.Prev()
	}
	return curNode
}

// SetPrev 设置上一个节点
func (me *Com) SetPrev(prev INode) {
	me.prev = prev
}

// Next 获取下一个节点
func (me *Com) Next() INode {
	return me.next
}

// SetNext 设置下一个节点
func (me *Com) SetNext(next INode) {
	me.next = next
}

// GetDef 根据名称获取定义
func (me *Com) GetDef(key string) INode {
	curNode := INode(me)
	// fmt.Printf("%v\n", curNode.SrcValue())
	for curNode != nil {
		fmt.Printf("%v\n", curNode.Type())
		if dict, ok := curNode.(*Dict); ok {
			dict.Print()
			if node, found := dict.DefNodeMap()[key]; found {
				return node
			}
		}
		curNode = curNode.Prev()
	}
	panic("获取不到目标定义")
}

// Type 类型
func (me *Com) Type() ENodeType {
	panic("node.Com: 抽象类被调用")
}

// BeginningOf 匹配
func (me *Com) BeginningOf(text string) (string, string, bool) {
	panic("node.Com: 抽象类被调用")
}
