package node

// End 结束节点
type End struct {
	prev INode
	*Com
}

// Type 节点类型
func (me *End) Type() ENodeType {
	return NodeTypeEnd
}

// BeginningOf s
func (me *End) BeginningOf(text string) (string, string) {
	return "", text
}

func (me *End) Next() INode {
	return nil
}

// NewEnd 构造函数
func NewEnd(prev INode) *End {
	return &End{
		prev: prev,
	}
}
