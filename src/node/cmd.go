package node

// Cmd 命令节点
type Cmd struct {
	cmd ENodeCmd
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

// BeginningOf 匹配
func (me *Cmd) BeginningOf(text string) (string, string, bool) {
	switch me.cmd {
	case NodeCmdOther:
		return "", text, true
	case NodeCmdEnd:
		return "", text, true
	default:
		panic("未知的命令")
	}
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
		cmd: getCmd(text),
		Com: Com{srcValue: text},
	}
}

// NewCmdEnd 构造结束节点
func NewCmdEnd() *Cmd {
	return &Cmd{
		cmd: NodeCmdEnd,
		Com: Com{srcValue: ".end"},
	}
}
