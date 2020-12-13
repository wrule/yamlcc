package node

// ENodeType 节点类型
type ENodeType int

const (
	// NodeTypeEnd 结束节点
	NodeTypeEnd ENodeType = iota
	// NodeTypeLit 字面文本节点
	NodeTypeLit
	// NodeTypeReg 正则表达式节点
	NodeTypeReg
	// NodeTypeDict 字典节点
	NodeTypeDict
	// NodeTypeDef 定义节点
	NodeTypeDef
	// NodeTypeRef 引用节点
	NodeTypeRef
	// NodeTypeBack 回跳节点
	NodeTypeBack
	// NodeTypeCmd 命令节点
	NodeTypeCmd
)
