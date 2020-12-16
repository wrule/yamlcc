package node

import "fmt"

// Cmd 命令节点
type Cmd struct {
	cmd ENodeCmd
	Com
}

// Cmd 命令字符串
func (me *Cmd) Cmd() ENodeCmd {
	return me.cmd
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
	rst := &Cmd{
		cmd: getCmd(text),
	}
	rst.Com = Com{me: rst, srcValue: text}
	return rst
}

// NewCmdEnd 构造结束节点
func NewCmdEnd() *Cmd {
	return NewCmd(".end")
}

// Print 打印节点信息
func (me *Cmd) Print() {
	fmt.Printf("命令节点: %v\n", me.SrcValue())
}
