package node

import "fmt"

// Def 定义节点
type Def struct {
	defName string
	Com
}

// DefName 定义名称
func (me *Def) DefName() string {
	return me.defName
}

// BeginningOf 匹配
func (me *Def) BeginningOf(text string) *Rst {
	return NewRst("", text, true)
}

// NewDef 构造函数
func NewDef(text string) *Def {
	rst := &Def{
		defName: text[1:],
	}
	rst.Com = NewCom(rst, text)
	return rst
}

// Print s
func (me *Def) Print() {
	fmt.Printf("定义节点: %v\n", me.DefName())
}
