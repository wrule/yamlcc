package node

// INode 接口
type INode interface {
	Me() INode
	Prev() INode
	PrevN(n int) INode
	SetPrev(INode)
	Nexts() []INode
	SetNexts([]INode)
	AppendNexts(INode)
	Link(INode)
	Links([]INode)

	SrcValue() interface{}
	IsEnd() bool
	IsNot() bool
	IsNextsEmpty() bool
	BeginningOf(string) (string, string, bool)
}
