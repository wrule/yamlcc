package node

// Other 其他命令节点
type Other struct {
	Com
}

// NewOther 构造函数
func NewOther() *Other {
	rst := &Other{}
	rst.Com = NewCom(rst, ".other")
	return rst
}
