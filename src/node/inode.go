package node

// INode 接口
type INode interface {
	Me() INode
	Prev() INode
	SetPrev(INode)
	Nexts() []INode
	SetNexts([]INode)
	AppendNexts(INode)

	SrcValue() interface{}
	IsEnd() bool
	BeginningOf(string) (string, string, bool)
}
