package main

type INode interface {
	Prev() INode
	PrevN(int) INode
	SetPrev(INode)
	Nexts() []INode
	Fasten(INode)
	Fastens([]INode)
	NextsIsEmpty() bool
	AppendNexts(INode)

	// 本节点头部匹配
	BeginningOf(string) *Rst

	IsRoot() bool
	IsReg() bool
	IsDef() bool
	IsRef() bool
	IsBack() bool
	IsEnd() bool
	IsOther() bool
	IsNot() bool

	// 节点自身链接
	Link()

	Print()
}
