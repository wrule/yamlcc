package main

import "fmt"

// Ref 引用节点
type Ref struct {
	defName string
	defNode *Def
	Com
}

// DefName 定义名称
func (me *Ref) DefName() string {
	return me.defName
}

// DefNode 定义节点
func (me *Ref) DefNode() *Def {
	return me.defNode
}

// Link 链接到定义
func (me *Ref) Link() {
	me.defNode = me.GetDef(me.DefName())
}

// BeginningOf s
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
