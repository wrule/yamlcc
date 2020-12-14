package node

// Back 回跳节点
type Back struct {
	value int
	prev  INode
	next  INode
	*Com
}

// Type 类型
func (me *Back) Type() ENodeType {
	return NodeTypeBack
}

// BeginningOf s
func (me *Back) BeginningOf(text string) (string, string) {
	return me.PrevN(me.value).BeginningOf(text)
}

// NewBack 构造函数
func NewBack(num int, prev INode) *Back {
	return &Back{
		value: num,
		prev:  prev,
		next:  nil,
	}
}
