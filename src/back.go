package main

import "fmt"

type Back struct {
	hops     int
	backNode INode
	Com
}

func (me *Back) Hops() int {
	return me.hops
}

func (me *Back) BackNode() INode {
	return me.backNode
}

func (me *Back) BeginningOf(text string) *Rst {
	return NewRst("", text, true)
}

func (me *Back) Link() {
	num := me.Hops() + 2
	if num >= 2 {
		me.backNode = me.PrevN(num)
	} else {
		me.backNode = me.PrevN(2)
	}
}

func NewBack(num int) *Back {
	rst := &Back{
		hops: num,
	}
	rst.Com = NewCom(rst, num)
	return rst
}

func (me *Back) Print() {
	fmt.Printf("回跳节点: %v\n", me.Hops())
}
