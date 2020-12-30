package node

import (
	"sort"
)

// BeginningOf s
func (me *Com) BeginningOf(text string) *Rst {
	return me.Me().BeginningOf(text)
}

// BeginningOfX s
func (me *Com) BeginningOfX(text string, trimHead bool) *Rst {
	ivdRst := NewRst("", text, true)
	// 按需进行头部无效字符修整匹配
	if trimHead {
		ivdRst = me.GetDef("invalid").NextsBeginningOfX(text, false)
	}
	// 进行本节点匹配
	meRst := me.Me().BeginningOf(ivdRst.Next())
	if meRst.Success() {
		// 进行子节点下推匹配
		nextRst := me.NextsBeginningOfX(meRst.Next(), true)
		if nextRst.Success() {
			return NewRst(ivdRst.Match()+meRst.Match()+nextRst.Match(), nextRst.Next(), true)
		}
		return NewRst(ivdRst.Match()+meRst.Match()+nextRst.Match(), nextRst.Next(), false)
	}
	return NewRst(ivdRst.Match()+meRst.Match(), meRst.Next(), false)
}

// NextsBeginningOfX s
func (me *Com) NextsBeginningOfX(text string, trimHead bool) *Rst {
	// 成功匹配结果列表
	successList := []*Rst{}
	// 失败匹配结果列表
	failureList := []*Rst{}
	// 遍历nextLogs匹配
	for _, node := range me.nextLogs {
		if rst := node.BeginningOfX(text, trimHead); rst.Success() {
			successList = append(successList, rst)
		} else {
			failureList = append(failureList, rst)
		}
	}

	// TODO: 另外这里也存在歧义逻辑需要后续开发处理

	// 如果成功结果数量为0，尝试执行.other逻辑
	if len(successList) < 1 && me.NextOther() != nil {
		if rst := me.NextOther().NextsBeginningOfX(text, trimHead); rst.Success() {
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
		if rst := not.BeginningOfX(text, true); rst.Success() {
			return false
		}
	}
	return true
}
