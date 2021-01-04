package main

import "fmt"

// Rst 匹配结果
type Rst struct {
	success bool
	match   string
	next    string
}

// Success 获取匹配状态
func (me *Rst) Success() bool {
	return me.success
}

// SetSuccess 更改匹配状态
func (me *Rst) SetSuccess(value bool) {
	me.success = value
}

// SuccessFull 获取完全匹配状态（完全匹配即完整匹配了整个字符串没有剩余）
func (me *Rst) SuccessFull() bool {
	return me.Success() && len(me.Next()) == 0
}

// Match 获取匹配文本
func (me *Rst) Match() string {
	return me.match
}

// Next 获取剩余文本
func (me *Rst) Next() string {
	return me.next
}

// Print 打印输出
func (me *Rst) Print() {
	fmt.Println("------------------")
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
