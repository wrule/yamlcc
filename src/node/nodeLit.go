package node

import "strings"

// Lit 字面文本节点
type Lit struct {
	value string
}

// Type s
func (me Lit) Type() ENodeType {
	return NodeTypeLit
}

// BeginningOf s
func (me *Lit) BeginningOf(text string) (string, string) {
	if strings.HasPrefix(text, me.value) {
		return me.value, text[len(me.value):]
	}
	return "", text
}

// NewLit 构造函数
func NewLit(text string) *Lit {
	return &Lit{
		value: text,
	}
}
