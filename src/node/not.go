package node

import "fmt"

// Not 其他命令
type Not struct {
	Com
}

// BeginningOf 匹配
func (me *Not) BeginningOf(text string) (string, string, bool) {
	return "", text, true
}

// NewNot 构造函数
func NewNot() *Not {
	rst := &Not{}
	rst.Com = Com{me: rst, srcValue: ".not"}
	return rst
}

// Print 打印节点信息
func (me *Not) Print() {
	fmt.Printf("非节点: %v\n", me.SrcValue())
}
