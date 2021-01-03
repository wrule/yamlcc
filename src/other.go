package main

import "fmt"

type Other struct {
	Com
}

func (me *Other) BeginningOf(text string) *Rst {
	return NewRst("", text, true)
}

func NewOther() *Other {
	rst := &Other{}
	rst.Com = NewCom(rst, nil)
	return rst
}

// Print s
func (me *Other) Print() {
	fmt.Printf("其他命令节点\n")
}
