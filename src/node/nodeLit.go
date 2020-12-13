package node

import "strings"

// Lit 字面文本节点
type Lit struct {
	value string
	prev  INode
	*Com
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

func (me *Lit) Next() INode {
	return NewEnd(me)
}

// NewLit 构造函数
func NewLit(text string, prev INode) *Lit {
	return &Lit{
		value: text,
		prev:  prev,
	}
}
