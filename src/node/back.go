package node

import "fmt"

// Back 回跳节点
type Back struct {
	level int
	Com
}

// Level 上跳层级
func (me *Back) Level() int {
	return me.level
}

// BeginningOf 匹配
func (me *Back) BeginningOf(text string) (string, string, bool) {
	return me.PrevDictN(me.level).BeginningOf(text)
}

// NewBack 构造函数
func NewBack(num int) *Back {
	rst := &Back{
		level: num,
	}
	rst.Com = Com{me: rst, srcValue: num}
	return rst
}

// Print 打印节点信息
func (me *Back) Print() {
	fmt.Printf("回跳节点: %v\n", me.SrcValue())
}
