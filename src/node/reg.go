package node

import (
	"fmt"
	"regexp"
)

// Reg 正则表达式节点
type Reg struct {
	regexp *regexp.Regexp
	Com
}

// Regexp 获取正则表达式
func (me *Reg) Regexp() *regexp.Regexp {
	return me.regexp
}

// BeginningOf 匹配
func (me *Reg) BeginningOf(text string) (string, string, bool) {
	indexs := me.regexp.FindStringIndex(text)
	if len(indexs) > 1 && indexs[0] == 0 {
		return text[:indexs[1]], text[indexs[1]:], true
	}
	return "", text, false
}

// NewReg 构造函数
func NewReg(text string) *Reg {
	rst := &Reg{
		regexp: regexp.MustCompile(text),
	}
	rst.Com = Com{me: rst, srcValue: text}
	return rst
}

// Print 打印节点信息
func (me *Reg) Print() {
	fmt.Printf("正则节点: %v\n", me.SrcValue())
}
