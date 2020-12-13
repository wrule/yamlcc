package node

// INode 节点接口
type INode interface {
	Type() ENodeType
	BeginningOf(string) (string, string)
	Prev() INode
	Next() INode
}
