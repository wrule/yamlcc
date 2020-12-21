package node

// Root 根节点
type Root struct {
	Com
}

// BeginningOf 匹配
func (me *Root) BeginningOf(text string) (string, string, bool) {
	return "", text, true
}

// NewRoot 构造函数
func NewRoot() *Root {
	rst := &Root{}
	rst.Com = NewCom(rst, nil)
	return rst
}
