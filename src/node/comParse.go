package node

import (
	"fmt"
	"sort"
)

// BeginningOf 节点头部匹配
func (me *Com) BeginningOf(text string) (string, string, bool) {
	panic("node.Com.BeginningOf: 抽象类方法被调用")
}

// BeginningTrimOf 节点头部修正匹配
func (me *Com) BeginningTrimOf(text string) *Rst {
	// 获取无效字符定义
	ivd := me.GetDef("invalid")
	// 进行头部无效字符匹配
	ivdRst := ivd.BeginningOf(text)

	fmt.Println("1111")
	ivdRst.Print()

	// 进行自身节点匹配
	meRst := me.Me().BeginningOf(ivdRst.Next())
	if meRst.Success() {
		nextRst := me.NextsBeginningTrimOf(meRst.Next())
		if nextRst.Success() {
			return NewRst(ivdRst.Match()+meRst.Match()+nextRst.Match(), nextRst.Next(), true)
		}
		return NewRst(ivdRst.Match()+meRst.Match()+nextRst.Match(), nextRst.Next(), false)
	}
	return NewRst(ivdRst.Match(), ivdRst.Next(), false)
}

// NextsBeginningTrimOf 节点Nexts的头部修正匹配
func (me *Com) NextsBeginningTrimOf(text string) *Rst {
	// 成功匹配结果列表
	successList := []*Rst{}
	// 失败匹配结果列表
	failureList := []*Rst{}
	// 遍历nextLogs匹配
	for _, node := range me.nextLogs {
		if rst := node.BeginningTrimOf(text); rst.Success() && me.NotsCheck(rst.Match()) {
			successList = append(successList, rst)
		} else {
			failureList = append(failureList, rst)
		}
	}

	// TODO: 另外这里也存在歧义逻辑需要后续开发处理

	// 如果成功结果数量为0，尝试执行.other逻辑
	if len(successList) < 1 && me.NextOther() != nil {
		if rst := me.NextOther().BeginningTrimOf(text); rst.Success() && me.NotsCheck(rst.Match()) {
			successList = append(successList, rst)
		}
	}

	// 对两个列表进行排序
	sort.Slice(successList, func(a, b int) bool {
		return len(successList[a].Match()) > len(successList[b].Match())
	})
	sort.Slice(failureList, func(a, b int) bool {
		return len(failureList[a].Match()) > len(failureList[b].Match())
	})
	// 排序后的匹配结果列表
	rstList := []*Rst{}
	rstList = append(rstList, successList...)
	rstList = append(rstList, failureList...)
	// 取贪心匹配最优结果返回
	if len(rstList) > 0 {
		return rstList[0]
	}
	return NewRst("", text, false)
}

// NotsCheck 非逻辑检查
func (me *Com) NotsCheck(text string) bool {
	for _, not := range me.nextNots {
		if rst := not.BeginningTrimOf(text); rst.Success() {
			return false
		}
	}
	return true
}
