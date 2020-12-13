package node

// Com 节点共用部分
type Com struct {
	prev INode
}

// Prev 获取父节点
func (me *Com) Prev() INode {
	return me.prev
}

// PrevN 获取n级的父节点（n大于等于1）
func (me *Com) PrevN(n int) INode {
	curNode := me.Prev()
	for i := 1; i < n; i++ {
		curNode = curNode.Prev()
	}
	return curNode
}

func (me *Com) GetDef(key string) INode {
	return NewLit("", nil)
}
