package node

// Prev s
func (me *Com) Prev() INode {
	return me.prev
}

// PrevN s
func (me *Com) PrevN(n int) INode {
	rst := me.Me()
	for i := 0; i < n && rst != nil; i++ {
		rst = rst.Prev()
	}
	if rst == nil {
		panic("node.Com.PrevN: 节点回跳越界")
	}
	return rst
}

// GetDef s
func (me *Com) GetDef(name string) *Def {
	curNode := me.Me()
	for curNode != nil {
		if def, found := curNode.NextDefs()[name]; found {
			return def
		}
		curNode = curNode.Prev()
	}
	panic("node.Com.GetDef: 找不到定义节点 " + name)
}

// SetPrev s
func (me *Com) SetPrev(prev INode) {
	me.prev = prev
}
