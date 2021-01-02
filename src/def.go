package main

import "fmt"

type Def struct {
	defName string
	Com
}

func (me *Def) DefName() string {
	return me.defName
}

func (me *Def) BeginningOf(text string) *Rst {
	return NewRst("", text, true)
}

func NewDef(text string) *Def {
	rst := &Def{
		defName: text[1:],
	}
	rst.Com = NewCom(rst, text)
	return rst
}

func (me *Def) Print() {
	fmt.Printf("定义节点: %v\n", me.DefName())
}
