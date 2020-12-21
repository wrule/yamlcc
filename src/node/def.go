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

// NewDef 构造函数
func NewDef(text string) *Def {
	rst := &Def{
		defName: text[1:],
	}
	rst.Com = NewCom(rst, text)
	return rst
}
