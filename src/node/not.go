package node

// Not 非命令节点
type Not struct {
	Com
}

// BeginningOf 匹配
func (me *Not) BeginningOf(text string) (string, string, bool) {
	return "", text, true
}

// NewNot 构造函数
func NewNot() *Not {
	rst := &Not{}
	rst.Com = NewCom(rst, ".not")
	return rst
}
