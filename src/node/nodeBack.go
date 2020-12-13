package node

// Back s
type Back struct {
	value int
	prev  INode
	*Com
}

// Type s
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
	}
}
