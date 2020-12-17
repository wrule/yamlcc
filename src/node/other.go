package node

import "fmt"

// Other 其他命令
type Other struct {
	Com
}

// BeginningOf 匹配
func (me *Other) BeginningOf(text string) (string, string, bool) {
	return "", text, true
}

// NewOther 构造函数
func NewOther() *Other {
	rst := &Other{}
	rst.Com = Com{me: rst, srcValue: ".other"}
	return rst
}

// Print 打印节点信息
func (me *Other) Print() {
	fmt.Printf("其他节点: %v\n", me.SrcValue())
}
