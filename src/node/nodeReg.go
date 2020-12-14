package node

import (
	"regexp"
)

// Reg 正则表达式节点
type Reg struct {
	value  string
	regexp *regexp.Regexp
	prev   INode
	next   INode
	*Com
}

// Regexp 获取正则表达式
func (me *Reg) Regexp() *regexp.Regexp {
	return me.regexp
}

// Type 类型
func (me *Reg) Type() ENodeType {
	return NodeTypeReg
}

// BeginningOf s
func (me *Reg) BeginningOf(text string) (string, string) {
	indexs := me.regexp.FindStringIndex(text)
	if len(indexs) > 1 && indexs[0] == 0 {
		return text[:indexs[1]], text[indexs[1]:]
	}
	return "", text
}

// NewReg 构造函数
func NewReg(text string) *Reg {
	return &Reg{
		value:  text,
		regexp: regexp.MustCompile(text),
	}
}
