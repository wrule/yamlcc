package main

import "fmt"

// Not 非命令节点
type Not struct {
	Com
}

// BeginningOf s
func (me *Not) BeginningOf(text string) *Rst {
	return NewRst("", text, true)
}

// NewNot 构造函数
func NewNot() *Not {
	rst := &Not{}
	rst.Com = NewCom(rst, ".not")
	return rst
}

// Print s
func (me *Not) Print() {
	fmt.Printf("非命令节点\n")
}
