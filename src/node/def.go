package node

// Def 定义节点
type Def struct {
	value   string
	defName string
	Com
}

// Type 类型
func (me Def) Type() ENodeType {
	return NodeTypeDef
}

// BeginningOf s
func (me *Def) BeginningOf(text string) (string, string) {
	return me.GetDef(me.defName).BeginningOf(text)
}

// NewDef 构造函数
func NewDef(text string) *Def {
	return &Def{
		value:   text,
		defName: text[1:],
	}
}
