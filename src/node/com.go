package node

import "sort"

// Com 公共逻辑（此结构相当于抽象类）
type Com struct {
	srcValue  interface{}
	me        INode
	prev      INode
	nexts     []INode
	nextDefs  map[string]*Def
	nextLogs  []INode
	nextCmds  []INode
	nextNots  []*Not
	nextOther *Other
}

// Me s
func (me *Com) Me() INode {
	return me.me
}

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

// Link s
func (me *Com) Link(next INode) {
	me.AppendNexts(next)
	next.SetPrev(me.Me())
}

// Links s
func (me *Com) Links(nexts []INode) {
	for _, next := range nexts {
		me.Link(next)
	}
}

// SrcValue s
func (me *Com) SrcValue() interface{} {
	return me.srcValue
}

// IsEnd s
func (me *Com) IsEnd() bool {
	_, ok := me.Me().(*End)
	return ok
}

// IsNot s
func (me *Com) IsNot() bool {
	_, ok := me.Me().(*Not)
	return ok
}

// IsOther s
func (me *Com) IsOther() bool {
	_, ok := me.Me().(*Other)
	return ok
}

// IsCmd s
func (me *Com) IsCmd() bool {
	return me.IsNot() || me.IsEnd() || me.IsOther()
}

// IsDef s
func (me *Com) IsDef() bool {
	_, ok := me.Me().(*Def)
	return ok
}

// IsLog s
func (me *Com) IsLog() bool {
	return !me.IsCmd() && !me.IsDef()
}

// IsNextsEmpty s
func (me *Com) IsNextsEmpty() bool {
	return len(me.Nexts()) < 1
}

// BeginningOf 节点头部匹配
func (me *Com) BeginningOf(text string) (string, string, bool) {
	panic("node.Com.BeginningOf: 抽象类方法被调用")
}

// BeginningTrimOf 节点头部修正匹配
func (me *Com) BeginningTrimOf(text string) *Rst {
	invalid := me.GetDef("invalid")
	// meMatch, meNext, meSuccess := me.Me().BeginningOf(ivdNext)
	return invalid.BeginningOf(text)
}

// NextsBeginningTrimOf 节点Nexts的头部修正匹配
func (me *Com) NextsBeginningTrimOf(text string) *Rst {
	// 成功匹配结果列表
	successList := []*Rst{}
	// 失败匹配结果列表
	failureList := []*Rst{}
	// 遍历nextLogs匹配
	for _, node := range me.nextLogs {
		if rst := node.BeginningTrimOf(text); rst.Success() && me.NotsCheck(text) {
			successList = append(successList, rst)
		} else {
			failureList = append(failureList, rst)
		}
	}
	// 这里还要考虑.other命令
	// 对两个列表进行排序
	sort.Slice(successList, func(a, b int) bool {
		return len(successList[a].Match()) > len(successList[b].Match())
	})
	sort.Slice(failureList, func(a, b int) bool {
		return len(failureList[a].Match()) > len(failureList[b].Match())
	})
	// 排序后的匹配结果列表
	rstList := []*Rst{}
	rstList = append(rstList, successList...)
	rstList = append(rstList, failureList...)
	if len(rstList) > 0 {
		return rstList[0]
	}
	return NewRst("", text, true)
}

// NotsCheck 非逻辑检查
func (me *Com) NotsCheck(text string) bool {
	for _, not := range me.nextNots {
		if rst := not.BeginningTrimOf(text); rst.Success() {
			return false
		}
	}
	return true
}

// NewCom 构造函数
func NewCom(me INode, srcValue interface{}) Com {
	return Com{
		me:       me,
		srcValue: srcValue,
	}
}
