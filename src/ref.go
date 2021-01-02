package main

import "fmt"

type Ref struct {
	defName string
	defNode *Def
	Com
}

func (me *Ref) DefName() string {
	return me.defName
}

func (me *Ref) DefNode() *Def {
	return me.defNode
}

func (me *Ref) BeginningOf(text string) *Rst {
	return NewRst("", text, true)
}

func NewRef(text string) *Ref {
	rst := &Ref{
		defName: text[1:],
	}
	rst.Com = NewCom(rst, text)
	return rst
}

func (me *Ref) Print() {
	fmt.Printf("引用节点: %v\n", me.DefName())
}
