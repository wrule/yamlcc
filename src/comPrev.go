package main

// Prev 获取上节点
func (me *Com) Prev() INode {
	return me.prev
}

// PrevN 获取上n个节点
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

// SetPrev 设置上节点
func (me *Com) SetPrev(prev INode) {
	me.prev = prev
}
