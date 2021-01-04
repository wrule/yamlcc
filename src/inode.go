package main

type INode interface {
	SrcValue() interface{}

	Prev() INode
	PrevN(int) INode
	SetPrev(INode)

	Nexts() []INode
	NextDefs() map[string]*Def
	NextLogs() []INode
	NextOther() *Other
	NextNots() []*Not
	NextsIsEmpty() bool

	Fasten(INode)
	Fastens([]INode)
	AppendNexts(INode)

	// 本节点头部匹配
	BeginningOf(string) *Rst
	BeginningOfX(string) *Rst
	NextsBeginningOfX(string) *Rst

	IsRoot() bool
	IsReg() bool
	IsDef() bool
	IsRef() bool
	IsBack() bool
	IsEnd() bool
	IsOther() bool
	IsNot() bool
	IsLog() bool

	// 节点初始化
	Init()
	// 节点自身链接
	Link()

	Print()
}
