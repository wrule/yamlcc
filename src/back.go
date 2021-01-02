package main

import "fmt"

type Back struct {
	hops int
	Com
}

func (me *Back) Hops() int {
	return me.hops
}

func (me *Back) BeginningOf(text string) *Rst {
	return NewRst("", text, true)
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
