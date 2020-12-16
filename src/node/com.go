package node

// Com 节点共用部分（抽象类）
type Com struct {
	srcValue interface{}
	prev     INode
	next     INode
	me       INode
}

// BeginningOf s
func (me *Com) BeginningOf(text string) (string, string, bool) {
	panic("node.Com.BeginningOf: 抽象类被调用")
}

// BeginningTrimOf s
func (me *Com) BeginningTrimOf(text string) (string, string, bool) {
	ivdMatch, ivdNext, _ := me.GetDefReg("invalid").BeginningOf(text)
	text = ivdNext
	iMe := me.Me()
	meMatch, meNext, meSuccess := iMe.BeginningOf(text)
	if me.IsEnd() {
		return meMatch, meNext, meSuccess
	}
	if meSuccess {
		meFullMatch := ivdMatch + meMatch
		nextMatch, nextNext, nextSuccess := me.Next().BeginningTrimOf(meNext)
		if nextSuccess {
			return meFullMatch + nextMatch, nextNext, nextSuccess
		}
		return meFullMatch, meNext, nextSuccess
	}
	return meMatch, meNext, meSuccess
}

// SrcValue s
func (me *Com) SrcValue() interface{} {
	return me.srcValue
}

// Me s
func (me *Com) Me() INode {
	return me.me
}

// IsEnd s
func (me *Com) IsEnd() bool {
	if cmd, ok := me.Me().(*Cmd); ok {
		if cmd.Cmd() == NodeCmdEnd {
			return true
		}
	}
	return false
}

// IsDict s
func (me *Com) IsDict() bool {
	_, ok := me.Me().(*Dict)
	return ok
}

// Prev s
func (me *Com) Prev() INode {
	return me.prev
}

// PrevN s
func (me *Com) PrevN(n int) INode {
	curNode := me.Me()
	for i := 0; i < n && curNode != nil; i++ {
		curNode = curNode.Prev()
	}
	if curNode == nil {
		panic("node.Com.PrevN: 目标上层为nil")
	}
	return curNode
}

// PrevDictN s
func (me *Com) PrevDictN(n int) *Dict {
	curNode := me.Me()
	for i := 0; curNode != nil; {
		if dict, ok := curNode.(*Dict); ok {
			i++
			if i >= n {
				return dict
			}
		}
		curNode = curNode.Prev()
	}
	panic("node.Com.PrevDictN: 目标上层为nil")
}

// SetPrev s
func (me *Com) SetPrev(prev INode) {
	me.prev = prev
}

// Next s
func (me *Com) Next() INode {
	return me.next
}

// SetNext s
func (me *Com) SetNext(next INode) {
	me.next = next
}

// GetDef s
func (me *Com) GetDef(key string) INode {
	curNode := me.Me()
	for curNode != nil {
		if dict, ok := curNode.(*Dict); ok {
			if node, found := dict.DefNodeMap()[key]; found {
				return node
			}
		}
		curNode = curNode.Prev()
	}
	panic("node.Com.GetDef: 获取不到定义")
}

// GetDefReg s
func (me *Com) GetDefReg(key string) *Reg {
	return me.GetDef(key).(*Reg)
}

// Print s
func (me *Com) Print() {
	panic("node.Com.Print: 抽象类被调用")
}
