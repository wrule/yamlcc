package node

import (
	"fmt"
	"sort"
)

// BeginningOf s
func (me *Com) BeginningOf(text string) *Rst {
	return me.Me().BeginningOf(text)
}

// BeginningOfX s
func (me *Com) BeginningOfX(text string) *Rst {
	meRst := me.Me().BeginningOf(text)
	if meRst.Success() {
		nextRst := me.NextsBeginningTrimOfX(meRst.Next())
		if nextRst.Success() {
			return NewRst(meRst.Match()+nextRst.Match(), nextRst.Next(), true)
		}
		return NewRst(meRst.Match()+nextRst.Match(), nextRst.Next(), false)
	}
	return meRst
}

// BeginningTrimOf s
func (me *Com) BeginningTrimOf(text string) *Rst {
	ivd := me.GetDef("invalid")
	// TODO
	ivdRst := ivd.BeginningOf(text)
	nodeRst := me.Me().BeginningOf(ivdRst.Next())
	return NewRst(ivdRst.Match()+nodeRst.Match(), nodeRst.Next(), nodeRst.Success())
}

// BeginningTrimOfX s
func (me *Com) BeginningTrimOfX(text string) *Rst {
	ivd := me.GetDef("invalid")
	ivdRst := ivd.BeginningOfX(text)
	fmt.Println("1111")
	ivdRst.Print()
	nodeRst := me.Me().BeginningOfX(ivdRst.Next())
	return NewRst(ivdRst.Match()+nodeRst.Match(), nodeRst.Next(), nodeRst.Success())
}

// NextsBeginningTrimOfX 节点Nexts的头部修正匹配
func (me *Com) NextsBeginningTrimOfX(text string) *Rst {
	// 成功匹配结果列表
	successList := []*Rst{}
	// 失败匹配结果列表
	failureList := []*Rst{}
	// 遍历nextLogs匹配
	for _, node := range me.nextLogs {
		if rst := node.BeginningTrimOfX(text); rst.Success() && me.NotsCheck(rst.Match()) {
			successList = append(successList, rst)
		} else {
			failureList = append(failureList, rst)
		}
	}

	// TODO: 另外这里也存在歧义逻辑需要后续开发处理

	// 如果成功结果数量为0，尝试执行.other逻辑
	if len(successList) < 1 && me.NextOther() != nil {
		if rst := me.NextOther().BeginningTrimOfX(text); rst.Success() && me.NotsCheck(rst.Match()) {
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

// NextsBeginningTrimOfX 节点Nexts的头部修正匹配
func (me *Com) NextsBeginningTrimOfX(text string) *Rst {
	// 成功匹配结果列表
	successList := []*Rst{}
	// 失败匹配结果列表
	failureList := []*Rst{}
	// 遍历nextLogs匹配
	for _, node := range me.nextLogs {
		if rst := node.BeginningTrimOfX(text); rst.Success() && me.NotsCheck(rst.Match()) {
			successList = append(successList, rst)
		} else {
			failureList = append(failureList, rst)
		}
	}

	// TODO: 另外这里也存在歧义逻辑需要后续开发处理

	// 如果成功结果数量为0，尝试执行.other逻辑
	if len(successList) < 1 && me.NextOther() != nil {
		if rst := me.NextOther().BeginningTrimOfX(text); rst.Success() && me.NotsCheck(rst.Match()) {
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
		if rst := not.BeginningTrimOfX(text); rst.Success() {
			return false
		}
	}
	return true
}
