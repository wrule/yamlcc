package main

import "fmt"

// Root 根节点
type Root struct {
	Com
}

// BeginningOf s
func (me *Root) BeginningOf(text string) *Rst {
	return NewRst("", text, true)
}

// NewRoot 构造函数
func NewRoot() *Root {
	rst := &Root{}
	rst.Com = NewCom(rst, ".root")
	return rst
}

// Print s
func (me *Root) Print() {
	fmt.Printf("根节点\n")
}
