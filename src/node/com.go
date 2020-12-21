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

// SrcValue s
func (me *Com) SrcValue() interface{} {
	return me.srcValue
}

// IsEnd s
func (me *Com) IsEnd() bool {
	return false
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
