package node

// Ref 引用节点
type Ref struct {
	value   string
	refName string
	prev    INode
	*Com
}

// Type s
func (me *Ref) Type() ENodeType {
	return NodeTypeRef
}

// BeginningOf s
func (me *Ref) BeginningOf(text string) (string, string) {
	return me.GetDef(me.refName).BeginningOf(text)
}

// NewRef 构造函数
func NewRef(text string, prev INode) *Ref {
	return &Ref{
		value:   text,
		refName: text[1:],
		prev:    prev,
	}
}
