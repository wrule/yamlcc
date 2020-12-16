package node

// Com 节点共用部分
type Com struct {
	srcValue interface{}
	prev     INode
	next     INode
	me       INode
}

// BeginningOf 匹配
func (me *Com) BeginningOf(text string) (string, string, bool) {
	panic("node.Com.BeginningOf: 抽象类被调用")
}

// IsEnd 判断节点是否是结束命令节点
func (me *Com) IsEnd() bool {
	if cmd, ok := me.Me().(*Cmd); ok {
		if cmd.Cmd() == NodeCmdEnd {
			return true
		}
	}
	return false
}

// BeginningTrimOf 字符串匹配（忽略无效字符）
func (me *Com) BeginningTrimOf(text string) (string, string, bool) {
	ivdMatch, ivdNext, _ := me.GetDefReg("invalid").BeginningOf(text)
	text = ivdNext
	iMe := me.Me()
	meMatch, meNext, meSuccess := iMe.BeginningOf(text)
	if me.IsEnd() {
		return meMatch, meNext, meSuccess
	}
	meFullMatch := ivdMatch + meMatch
	if meSuccess {
		nextMatch, nextNext, nextSuccess := me.Next().BeginningTrimOf(meNext)
		if nextSuccess {
			return meFullMatch + nextMatch, nextNext, true
		}
		return meFullMatch, meNext, false
	}
	return "", text, false
}

// SrcValue 原始值
func (me *Com) SrcValue() interface{} {
	return me.srcValue
}

// Me 节点实例（非抽象类）
func (me *Com) Me() INode {
	return me.me
}

// Prev 获取上一个节点
func (me *Com) Prev() INode {
	return me.prev
}

// PrevN 获取上n个节点
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

// SetPrev 设置上一个节点
func (me *Com) SetPrev(prev INode) {
	me.prev = prev
}

// Next 获取下一个节点
func (me *Com) Next() INode {
	return me.next
}

// SetNext 设置下一个节点
func (me *Com) SetNext(next INode) {
	me.next = next
}

// GetDef 根据名称获取定义节点
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

// GetDefReg 根据名称获取正则节点
func (me *Com) GetDefReg(key string) *Reg {
	return me.GetDef(key).(*Reg)
}

// Print 打印信息
func (me *Com) Print() {
	panic("node.Com.Print: 抽象类被调用")
}
