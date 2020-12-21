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

// NewRef 构造函数
func NewRef(text string) *Ref {
	rst := &Ref{
		defName: text[1:],
	}
	rst.Com = NewCom(rst, text)
	return rst
}
