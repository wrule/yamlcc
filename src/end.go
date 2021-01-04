package main

import "fmt"

// End 结束命令节点
type End struct {
	Com
}

// BeginningOf s
func (me *End) BeginningOf(text string) *Rst {
	return NewRst("", text, true)
}

// NewEnd 构造函数
func NewEnd() *End {
	rst := &End{}
	rst.Com = NewCom(rst, ".end")
	return rst
}

// Print s
func (me *End) Print() {
	fmt.Printf("结束命令节点\n")
}
