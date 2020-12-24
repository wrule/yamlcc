package node

import "fmt"

// End 结束命令节点
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
	rst.Com = NewCom(rst, ".end")
	return rst
}

// Print s
func (me *End) Print() {
	fmt.Printf("结束命令节点\n")
}
