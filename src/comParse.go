package main

import (
	"sort"
)

// BeginningOf s
func (me *Com) BeginningOf(text string) *Rst {
	return me.Me().BeginningOf(text)
}

// BeginningOfX s
func (me *Com) BeginningOfX(text string) *Rst {
	rst := me.BeginningOf(text)
	if rst.Success() {
		nextRst := me.NextsBeginningOfX(rst.Next())
		if nextRst.Success() {
			return NewRst(rst.Match()+nextRst.Match(), nextRst.Next(), true)
		}
		return NewRst(rst.Match()+nextRst.Match(), nextRst.Next(), false)
	}
	return rst
}

// NextsBeginningOfX s
func (me *Com) NextsBeginningOfX(text string) *Rst {
	successList := []*Rst{}
	failureList := []*Rst{}
	for _, log := range me.NextLogs() {
		rst := log.BeginningOfX(text)
		if rst.Success() && me.NotsCheck(rst.Match()) {
			successList = append(successList, rst)
		} else {
			rst.SetSuccess(false)
			failureList = append(failureList, rst)
		}
	}

	// .other逻辑
	if len(successList) < 1 && me.NextOther() != nil {
		if rst := me.NextOther().BeginningOfX(text); rst.Success() {
			successList = append(successList, rst)
		} else {
			failureList = append(failureList, rst)
		}
	}

	sort.Slice(successList, func(a, b int) bool {
		return len(successList[a].Match()) > len(successList[b].Match())
	})
	sort.Slice(failureList, func(a, b int) bool {
		return len(failureList[a].Match()) > len(failureList[b].Match())
	})
	rstList := []*Rst{}
	rstList = append(rstList, successList...)
	rstList = append(rstList, failureList...)
	if len(rstList) > 0 {
		return rstList[0]
	}
	return NewRst("", text, true)
}

// NotsCheck s
func (me *Com) NotsCheck(text string) bool {
	for _, not := range me.NextNots() {
		if rst := not.BeginningOfX(text); rst.SuccessFull() {
			return false
		}
	}
	return true
}
