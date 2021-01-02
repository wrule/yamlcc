package node

// Com 公共逻辑（此结构相当于抽象类）
type Com struct {
	srcValue  interface{}
	me        INode
	prev      INode
	nexts     []INode
	nextDefs  map[string]*Def
	nextLogs  []INode
	nextNots  []*Not
	nextOther *Other
	inInvalid bool
}

// Init 初始化
func (me *Com) Init() {
	if me.SrcValue() == ":invalid" {
		me.inInvalid = true
	} else if me.Prev().InInvalid() {
		me.inInvalid = true
	}
}

// InInvalid s
func (me *Com) InInvalid() bool {
	return me.inInvalid
}

// SrcValue s
func (me *Com) SrcValue() interface{} {
	return me.srcValue
}

// Me s
func (me *Com) Me() INode {
	return me.me
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
	return me.IsEnd() || me.IsNot() || me.IsOther()
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

// IsNextLogsEmpty s
func (me *Com) IsNextLogsEmpty() bool {
	return len(me.NextLogs()) < 1
}

// NewCom 构造函数
func NewCom(me INode, srcValue interface{}) Com {
	return Com{
		me:       me,
		srcValue: srcValue,
	}
}
