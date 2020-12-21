package node

// Other 其他命令节点
type Other struct {
	Com
}

// BeginningOf 匹配
func (me *Other) BeginningOf(text string) (string, string, bool) {
	return "", text, true
}

// NewOther 构造函数
func NewOther() *Other {
	rst := &Other{}
	rst.Com = NewCom(rst, ".other")
	return rst
}
