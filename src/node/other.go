package node

import "fmt"

// Other 其他命令节点
type Other struct {
	Com
}

// BeginningOf 匹配
func (me *Other) BeginningOf(text string) *Rst {
	return me.NextsBeginningOfX(text, false)
}

// NewOther 构造函数
func NewOther() *Other {
	rst := &Other{}
	rst.Com = NewCom(rst, ".other")
	return rst
}

// Print s
func (me *Other) Print() {
	fmt.Printf("其他命令节点\n")
}
