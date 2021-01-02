package main

type INode interface {
	Prev() INode
	PrevN(int) INode
	SetPrev(INode)
	Link(INode)
	Links([]INode)
	NextsIsEmpty() bool

	// 本节点头部匹配
	BeginningOf(string) *Rst

	Print()
}
