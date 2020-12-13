package node

// INode 节点接口
type INode interface {
	// 节点类型
	Type() ENodeType
	// 上一个节点
	Prev() INode
	// 下一个节点
	Next() INode
	// 字符串匹配
	BeginningOf(string) (string, string)
}
