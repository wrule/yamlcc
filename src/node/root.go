package node

import "fmt"

// Root 根节点
type Root struct {
	Com
}

// BeginningOf 匹配
func (me *Root) BeginningOf(text string) *Rst {
	return me.NextsBeginningOfX(text, false)
}

// NewRoot 构造函数
func NewRoot() *Root {
	rst := &Root{}
	rst.Com = NewCom(rst, nil)
	return rst
}

// Print s
func (me *Root) Print() {
	fmt.Printf("根节点\n")
}
