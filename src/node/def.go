package node

// Def 定义节点
type Def struct {
	defName string
	Com
}

// DefName 定义名称
func (me *Def) DefName() string {
	return me.defName
}

// Type 类型
func (me *Def) Type() ENodeType {
	return NodeTypeDef
}

// BeginningOf 匹配
func (me *Def) BeginningOf(text string) (string, string, bool) {
	return "", text, false
}

// NewDef 构造函数
func NewDef(text string) *Def {
	return &Def{
		defName: text[1:],
		Com:     Com{srcValue: text},
	}
}
