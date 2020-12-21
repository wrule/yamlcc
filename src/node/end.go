package node

// End 结束命令节点
type End struct {
	Com
}

// NewEnd 构造函数
func NewEnd() *End {
	rst := &End{}
	rst.Com = NewCom(rst, ".end")
	return rst
}
