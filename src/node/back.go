package node

// Back 回跳节点
type Back struct {
	level int
	Com
}

// Level 上跳层级
func (me *Back) Level() int {
	return me.level
}

// Type 类型
func (me *Back) Type() ENodeType {
	return NodeTypeBack
}

// BeginningOf 匹配
func (me *Back) BeginningOf(text string) (string, string, bool) {
	return me.PrevN(me.level).BeginningOf(text)
}

// NewBack 构造函数
func NewBack(num int) *Back {
	rst := &Back{
		level: num,
	}
	rst.Com = Com{me: rst, srcValue: num}
	return rst
}
