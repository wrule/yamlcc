package main

type INode interface {
	Prev() INode
	PrevN(int) INode
	SetPrev(INode)
	Fasten(INode)
	Fastens([]INode)
	NextsIsEmpty() bool

	// 本节点头部匹配
	BeginningOf(string) *Rst

	Print()
}
