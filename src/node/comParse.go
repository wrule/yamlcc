package node

import "sort"

// BeginningOf 节点头部匹配
func (me *Com) BeginningOf(text string) (string, string, bool) {
	panic("node.Com.BeginningOf: 抽象类方法被调用")
}

// BeginningTrimOf 节点头部修正匹配
func (me *Com) BeginningTrimOf(text string) *Rst {
	invalid := me.GetDef("invalid")
	// meMatch, meNext, meSuccess := me.Me().BeginningOf(ivdNext)
	return invalid.BeginningOf(text)
}

// NextsBeginningTrimOf 节点Nexts的头部修正匹配
func (me *Com) NextsBeginningTrimOf(text string) *Rst {
	// 成功匹配结果列表
	successList := []*Rst{}
	// 失败匹配结果列表
	failureList := []*Rst{}
	// 遍历nextLogs匹配
	for _, node := range me.nextLogs {
		if rst := node.BeginningTrimOf(text); rst.Success() && me.NotsCheck(text) {
			successList = append(successList, rst)
		} else {
			failureList = append(failureList, rst)
		}
	}
	// 这里还要考虑.other命令
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
	if len(rstList) > 0 {
		return rstList[0]
	}
	return NewRst("", text, true)
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
