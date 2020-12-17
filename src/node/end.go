package node

import "fmt"

// End 结束命令
type End struct {
	Com
}

// BeginningOf 匹配
func (me *End) BeginningOf(text string) (string, string, bool) {
	return "", text, true
}

// NewEnd 构造函数
func NewEnd() *End {
	rst := &End{}
	rst.Com = Com{me: rst, srcValue: ".end"}
	return rst
}

// Print 打印节点信息
func (me *End) Print() {
	fmt.Printf("结束节点: %v\n", me.SrcValue())
}
