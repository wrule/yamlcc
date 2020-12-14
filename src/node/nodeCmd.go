package node

// Cmd 命令节点
type Cmd struct {
	value string
	cmd   ENodeCmd
	Com
}

// Cmd 命令字符串
func (me *Cmd) Cmd() ENodeCmd {
	return me.cmd
}

// Type 节点类型
func (me *Cmd) Type() ENodeType {
	return NodeTypeCmd
}

// BeginningOf s
func (me *Cmd) BeginningOf(text string) (string, string) {
	return "", text
}

func getCmd(text string) ENodeCmd {
	switch text[1:] {
	case "end":
		return NodeCmdEnd
	case "other":
		return NodeCmdOther
	}
	panic("未知的命令")
}

// NewCmd 构造函数
func NewCmd(text string) *Cmd {
	return &Cmd{
		value: text,
		cmd:   getCmd(text),
	}
}

// NewCmdEnd s
func NewCmdEnd() *Cmd {
	return &Cmd{
		value: ".end",
		cmd:   NodeCmdEnd,
	}
}
