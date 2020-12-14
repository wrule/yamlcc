package node

// ENodeCmd 命令节点类型
type ENodeCmd int

const (
	// NodeCmdEnd 结束命令
	NodeCmdEnd ENodeCmd = iota
	// NodeCmdOther Other命令
	NodeCmdOther
)
