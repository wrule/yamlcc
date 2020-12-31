package node

import "fmt"

// Ref 引用节点
type Ref struct {
	defName string
	defNode INode
	Com
}

// DefName 定义名称
func (me *Ref) DefName() string {
	return me.defName
}

// DefNode 定义节点
func (me *Ref) DefNode() INode {
	return me.defNode
}

// RefLink 链接定义
func (me *Ref) RefLink() {
	me.defNode = me.GetDef(me.DefName())
}

// BeginningOf 匹配
func (me *Ref) BeginningOf(text string) *Rst {
	return me.DefNode().BeginningOf(text)
}

// NewRef 构造函数
func NewRef(text string) *Ref {
	rst := &Ref{
		defName: text[1:],
	}
	rst.Com = NewCom(rst, text)
	return rst
}

// Print s
func (me *Ref) Print() {
	fmt.Printf("引用节点: %v\n", me.DefName())
}
