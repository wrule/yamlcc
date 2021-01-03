package main

import "fmt"

type Root struct {
	Com
}

func (me *Root) BeginningOf(text string) *Rst {
	return NewRst("", text, true)
}

func NewRoot() *Root {
	rst := &Root{}
	rst.Com = NewCom(rst, ".root")
	return rst
}

func (me *Root) Print() {
	fmt.Printf("根节点\n")
}
