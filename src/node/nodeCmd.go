package node

// Cmd 命令节点
type Cmd struct {
	value string
	prev  INode
	*Com
}

// Type 节点类型
func (me *Cmd) Type() ENodeType {
	return NodeTypeCmd
}

// BeginningOf s
func (me *Cmd) BeginningOf(text string) (string, string) {
	return "", text
}

// NewCmd 构造函数
func NewCmd(text string, prev INode) *Cmd {
	return &Cmd{
		value: text,
		prev:  prev,
	}
}
