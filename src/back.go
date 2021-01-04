package main

import "fmt"

// Back 回跳节点
type Back struct {
	hops     int
	backNode INode
	Com
}

// Hops 跳数
func (me *Back) Hops() int {
	return me.hops
}

// BackNode 回跳目标节点
func (me *Back) BackNode() INode {
	return me.backNode
}

// BeginningOf s
func (me *Back) BeginningOf(text string) *Rst {
	return me.BackNode().NextsBeginningOfX(text)
}

// Link 链接到回跳
func (me *Back) Link() {
	num := me.Hops() + 2
	if num >= 2 {
		me.backNode = me.PrevN(num)
	} else {
		me.backNode = me.PrevN(2)
	}
}

// NewBack 构造函数
func NewBack(num int) *Back {
	rst := &Back{
		hops: num,
	}
	rst.Com = NewCom(rst, num)
	return rst
}

// Print s
func (me *Back) Print() {
	fmt.Printf("回跳节点: %v\n", me.SrcValue())
}
