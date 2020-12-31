package node

import "fmt"

// Back 回跳节点
type Back struct {
	hops int
	Com
}

// Hops 跳数
func (me *Back) Hops() int {
	return me.hops
}

// BackNode 回跳目标节点
func (me *Back) BackNode() INode {
	num := me.Hops() + 2
	if num >= 2 {
		return me.PrevN(num)
	}
	return me.PrevN(2)
}

// BeginningOf s
func (me *Back) BeginningOf(text string) *Rst {
	return me.PrevN(me.hops).BeginningOf(text)
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
	fmt.Printf("回跳命令节点: %v\n", me.Hops())
}
