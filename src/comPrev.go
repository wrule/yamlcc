package main

func (me *Com) Prev() INode {
	return me.prev
}

func (me *Com) PrevN(n int) INode {
	rst := me.Me()
	for i := 0; i < n && rst != nil; i++ {
		rst = rst.Prev()
	}
	if rst == nil {
		panic("main.Com.PrevN: 回跳越界")
	}
	return rst
}

func (me *Com) SetPrev(prev INode) {
	me.prev = prev
}
