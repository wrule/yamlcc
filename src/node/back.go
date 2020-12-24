package node

// Back 回跳节点
type Back struct {
	hops int
	Com
}

// Hops 跳数
func (me *Back) Hops() int {
	return me.hops
}

// BeginningOf s
func (me *Back) BeginningOf(text string) (string, string, bool) {
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
