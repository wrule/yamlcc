package node

import "sort"

// BeginningOf s
func (me *Com) BeginningOf(text string) *Rst {
	return me.Me().BeginningOf(text)
}

// BeginningOfX s
func (me *Com) BeginningOfX(text string, trimHead bool) *Rst {
	// 按需进行头部修整匹配
	ivdRst := me.invalidTrimHead(text, trimHead)
	// 进行节点的原生匹配
	meRst := me.Me().BeginningOf(ivdRst.Next())
	if meRst.Success() {
		// 进行下节点的下推匹配（头部修整）
		nextRst := me.Me().NextsBeginningOfX(meRst.Next(), true)
		if nextRst.Success() {
			return NewRst(ivdRst.Match()+meRst.Match()+nextRst.Match(), nextRst.Next(), true)
		}
		return NewRst(ivdRst.Match()+meRst.Match()+nextRst.Match(), nextRst.Next(), false)
	}
	return NewRst(ivdRst.Match()+meRst.Match(), meRst.Next(), false)
}

// NextsBeginningOfX s
func (me *Com) NextsBeginningOfX(text string, trimHead bool) *Rst {
	// 默认返回值（根据trimHead决定的修整匹配结果，且success为false）
	rst := me.invalidTrimHead(text, trimHead)
	rst.SetSuccess(false)
	// 非逻辑检查
	if me.NotsCheck(text, trimHead) {
		// 定义成功失败结果列表
		successList := []*Rst{}
		failureList := []*Rst{}
		// 遍历下推匹配逻辑节点，并且采集成功失败结果
		for _, log := range me.Me().NextLogs() {
			rst := log.BeginningOfX(text, trimHead)
			if rst.Success() {
				successList = append(successList, rst)
			} else {
				failureList = append(failureList, rst)
			}
		}
		// 其他命令逻辑
		if len(successList) < 1 && me.Me().NextOther() != nil {
			if rst := me.Me().NextOther().BeginningOfX(text, trimHead); rst.Success() {
				successList = append(successList, rst)
			} else {
				failureList = append(failureList, rst)
			}
		}
		// 按匹配长度排序（贪心）（这里有可能存在歧义要处理哦）
		sort.Slice(successList, func(a, b int) bool {
			return len(successList[a].Match()) > len(successList[b].Match())
		})
		sort.Slice(failureList, func(a, b int) bool {
			return len(failureList[a].Match()) > len(failureList[b].Match())
		})
		// 拼接结果
		rstList := []*Rst{}
		rstList = append(rstList, successList...)
		rstList = append(rstList, failureList...)
		// 返回结果
		if len(rstList) > 0 {
			return rstList[0]
		}
	}
	return rst
}

func (me *Com) invalidTrimHead(text string, trimHead bool) *Rst {
	ivdRst := NewRst("", text, true)
	if trimHead {
		ivdRst = me.GetInvalid().BeginningOf(text)
	}
	return ivdRst
}

// NotsCheck 非逻辑检查
func (me *Com) NotsCheck(text string, trimHead bool) bool {
	for _, not := range me.Me().NextNots() {
		if rst := not.Me().BeginningOfX(text, trimHead); rst.Success() {
			return false
		}
	}
	return true
}
