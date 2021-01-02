package main

type INode interface {
	SetPrev(INode)
	Link(INode)
	Links([]INode)
	NextsIsEmpty() bool

	// 本节点头部匹配
	BeginningOf(string) *Rst

	Print()
}
