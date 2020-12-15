package node

// Ref 引用节点
type Ref struct {
	defName string
	Com
}

// DefName 定义名称
func (me *Ref) DefName() string {
	return me.defName
}

// Type 类型
func (me *Ref) Type() ENodeType {
	return NodeTypeRef
}

// BeginningOf 匹配
func (me *Ref) BeginningOf(text string) (string, string, bool) {
	return me.GetDef(me.defName).BeginningOf(text)
}

// NewRef 构造函数
func NewRef(text string) *Ref {
	return &Ref{
		defName: text[1:],
		Com:     Com{srcValue: text},
	}
}
