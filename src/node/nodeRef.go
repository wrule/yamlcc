package node

// Ref 引用节点
type Ref struct {
	value   string
	refName string
	Com
}

// Type 类型
func (me *Ref) Type() ENodeType {
	return NodeTypeRef
}

// BeginningOf s
func (me *Ref) BeginningOf(text string) (string, string) {
	return me.GetDef(me.refName).BeginningOf(text)
}

// NewRef 构造函数
func NewRef(text string) *Ref {
	return &Ref{
		value:   text,
		refName: text[1:],
	}
}
