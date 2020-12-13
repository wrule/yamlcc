package node

import (
	"regexp"
)

// Reg 正则表达式节点
type Reg struct {
	value string
	re    *regexp.Regexp
	prev  INode
}

// Type s
func (me Reg) Type() ENodeType {
	return NodeTypeReg
}

// BeginningOf s
func (me *Reg) BeginningOf(text string) (string, string) {
	indexs := me.re.FindStringIndex(text)
	if len(indexs) > 1 && indexs[0] == 0 {
		return text[:indexs[1]], text[indexs[1]:]
	}
	return "", text
}

// NewReg 构造函数
func NewReg(text string, prev INode) *Reg {
	return &Reg{
		value: text,
		re:    regexp.MustCompile(text),
		prev:  prev,
	}
}
