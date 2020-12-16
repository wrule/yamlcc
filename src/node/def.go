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

// BeginningOf 匹配
func (me *Def) BeginningOf(text string) (string, string, bool) {
	return "", text, false
}

// NewDef 构造函数
func NewDef(text string) *Def {
	rst := &Def{
		defName: text[1:],
	}
	rst.Com = Com{me: rst, srcValue: text}
	return rst
}
