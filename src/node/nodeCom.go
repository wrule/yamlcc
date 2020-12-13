package node

// Com 节点共用部分
type Com struct {
	prev INode
	next INode
}

// Prev 获取上一个节点
func (me *Com) Prev() INode {
	return me.prev
}

// PrevN 获取上n个节点（n大于等于1）
func (me *Com) PrevN(n int) INode {
	curNode := me.Prev()
	for i := 1; i < n; i++ {
		curNode = curNode.Prev()
	}
	return curNode
}

// Next 获取下一个节点
func (me *Com) Next() INode {
	return me.next
}

// GetDef 获取定义
func (me *Com) GetDef(key string) INode {
	return NewLit("", nil, nil)
}
