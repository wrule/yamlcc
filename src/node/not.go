package node

// Not 非命令节点
type Not struct {
	Com
}

// NewNot 构造函数
func NewNot() *Not {
	rst := &Not{}
	rst.Com = NewCom(rst, ".not")
	return rst
}
