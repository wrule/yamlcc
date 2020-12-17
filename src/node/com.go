package node

import "fmt"

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
	// 修剪文本头部无效字符
	ivdMatch, textTrimmed, _ := me.GetDefReg("invalid").BeginningOf(text)
	// 进行非Trim匹配
	meMatch, meNext, meSuccess := me.Me().BeginningOf(textTrimmed)
	meFullMatch := ivdMatch + meMatch
	if me.IsEnd() {
		return meMatch, text, meSuccess
	}
	if meSuccess {
		nextMatch, nextNext, nextSuccess := me.Next().BeginningTrimOf(meNext)
		if nextSuccess {
			return meFullMatch + nextMatch, nextNext, nextSuccess
		}
		return meFullMatch, meNext, nextSuccess
	}
	return meMatch, text, meSuccess
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
	_, ok := me.Me().(*End)
	return ok
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
			if i > n {
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
	fmt.Println(key)
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
