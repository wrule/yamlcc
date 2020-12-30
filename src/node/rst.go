package node

import "fmt"

// Rst 匹配结果
type Rst struct {
	success bool
	match   string
	next    string
}

// Success 匹配是否成功
func (me *Rst) Success() bool {
	return me.success
}

// Match 匹配获得文本
func (me *Rst) Match() string {
	return me.match
}

// Next 匹配剩余文本
func (me *Rst) Next() string {
	return me.next
}

// Print s
func (me *Rst) Print() {
	fmt.Printf("success: %v\n", me.Success())
	fmt.Printf("match(%d):\n[%s]\n", len(me.Match()), me.Match())
	fmt.Printf("next(%d):\n[%s]\n", len(me.Next()), me.Next())
}

// NewRst 构造函数
func NewRst(match, next string, success bool) *Rst {
	return &Rst{
		success: success,
		match:   match,
		next:    next,
	}
}
