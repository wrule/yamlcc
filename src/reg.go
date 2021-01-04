package main

import (
	"fmt"
	"regexp"
)

// Reg 正则节点
type Reg struct {
	re *regexp.Regexp
	Com
}

// Regexp 获取正则表达式
func (me *Reg) Regexp() *regexp.Regexp {
	return me.re
}

// BeginningOf s
func (me *Reg) BeginningOf(text string) *Rst {
	indexs := me.Regexp().FindStringIndex(text)
	if len(indexs) > 1 && indexs[0] == 0 {
		return NewRst(text[:indexs[1]], text[indexs[1]:], true)
	}
	return NewRst("", text, false)
}

// NewReg 构造函数
func NewReg(text string) *Reg {
	rst := &Reg{
		re: regexp.MustCompile(text),
	}
	rst.Com = NewCom(rst, text)
	return rst
}

// Print s
func (me *Reg) Print() {
	fmt.Printf("正则节点: %v\n", me.SrcValue())
}
