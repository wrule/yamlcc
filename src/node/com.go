package node

// Com 公共逻辑（此结构相当于抽象类）
type Com struct {
	srcValue interface{}
	me       INode
	prev     INode
	nexts    []INode
	nextDefs map[string]*Def
	nextLogs []INode
	nextCmds []INode
	nextNots []*Not
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

// Nexts s
func (me *Com) Nexts() []INode {
	return me.nextLogs
}

// NextDefs s
func (me *Com) NextDefs() map[string]*Def {
	return me.nextDefs
}

// NextCmds s
func (me *Com) NextCmds() []INode {
	return me.nextCmds
}

// NextNots s
func (me *Com) NextNots() []*Not {
	return me.nextNots
}

// SetNexts s
func (me *Com) SetNexts(nexts []INode) {
	me.nexts = nexts
	me.updateNexts()
}

// AppendNexts s
func (me *Com) AppendNexts(next INode) {
	me.nexts = append(me.nexts, next)
	me.updateNexts()
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

// BeginningOf s
func (me *Com) BeginningOf(text string) (string, string, bool) {
	panic("node.Com.BeginningOf: 抽象类方法被调用")
}

// BeginningTrimOf s
func (me *Com) BeginningTrimOf(text string) (string, string, bool) {
	invalid := me.GetDef("invalid")
	ivdMatch, ivdNext, ivdSuccess := invalid.BeginningOf(text)
	// meMatch, meNext, meSuccess := me.Me().BeginningOf(ivdNext)
	return ivdMatch, ivdNext, ivdSuccess
}

func (me *Com) NextBeginningTrimOf(text string) (string, string, bool) {
	// for _, node := range me.nextLogs {
	// curMatch, curNext, curSuccess := node.BeginningTrimOf(text)
	// }
	return "", "", true
}

// updateNextDefs 同步更新nextDefs
func (me *Com) updateNextDefs() {
	me.nextDefs = map[string]*Def{}
	for _, node := range me.nexts {
		if node.IsDef() {
			def := node.(*Def)
			me.nextDefs[def.DefName()] = def
		}
	}
}

// updateNextLogs 同步更新nextLogs
func (me *Com) updateNextLogs() {
	me.nextLogs = []INode{}
	for _, node := range me.nexts {
		if node.IsLog() {
			me.nextLogs = append(me.nextLogs, node)
		}
	}
}

// updateNextCmds 同步更新nextCmds
func (me *Com) updateNextCmds() {
	me.nextCmds = []INode{}
	for _, node := range me.nexts {
		if node.IsCmd() {
			me.nextCmds = append(me.nextCmds, node)
		}
	}
}

// updateNextNots 同步更新nextNots
func (me *Com) updateNextNots() {
	me.nextNots = []*Not{}
	for _, node := range me.nexts {
		if node.IsNot() {
			not := node.(*Not)
			me.nextNots = append(me.nextNots, not)
		}
	}
}

// updateNexts 同步更新其他的nexts相关数据结构
func (me *Com) updateNexts() {
	me.updateNextDefs()
	me.updateNextLogs()
	me.updateNextCmds()
	me.updateNextNots()
}

// NewCom 构造函数
func NewCom(me INode, srcValue interface{}) Com {
	return Com{
		me:       me,
		srcValue: srcValue,
	}
}
