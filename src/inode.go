package node

type INode interface {
	// 本节点头部匹配
	BeginningOf(string) *Rst

	Print()
}
