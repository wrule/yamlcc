package node

import "regexp"

// Reg 正则节点
type Reg struct {
	re *regexp.Regexp
	Com
}

// Regexp 获取正则表达式
func (me *Reg) Regexp() *regexp.Regexp {
	return me.re
}

// NewReg 构造函数
func NewReg(text string) *Reg {
	rst := &Reg{
		re: regexp.MustCompile(text),
	}
	rst.Com = NewCom(rst, text)
	return rst
}
