package node

// Com 公共逻辑
type Com struct {
	me       INode
	prev     INode
	nexts    []INode
	srcValue interface{}
}

// Me s
func (me *Com) Me() INode {
	return me.me
}

// Prev s
func (me *Com) Prev() INode {
	return me.prev
}

// SetPrev s
func (me *Com) SetPrev(prev INode) {
	me.prev = prev
}

// Nexts s
func (me *Com) Nexts() []INode {
	return me.nexts
}

// SetNexts s
func (me *Com) SetNexts(nexts []INode) {
	me.nexts = nexts
}

// AppendNexts s
func (me *Com) AppendNexts(next INode) {
	me.nexts = append(me.nexts, next)
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

// IsNextsEmpty s
func (me *Com) IsNextsEmpty() bool {
	return len(me.Nexts()) < 1
}

// BeginningOf s
func (me *Com) BeginningOf(text string) (string, string, bool) {
	panic("node.Com.BeginningOf: 抽象类方法被调用")
}

// NewCom 构造函数
func NewCom(me INode, srcValue interface{}) Com {
	return Com{
		me:       me,
		srcValue: srcValue,
	}
}
