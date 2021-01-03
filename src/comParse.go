package main

import (
	"sort"
)

func (me *Com) BeginningOfX(text string) *Rst {
	meRst := me.Me().BeginningOf(text)
	if meRst.Success() {
		nextRst := me.Me().NextsBeginningOfX(meRst.Next())
		if nextRst.Success() {
			return NewRst(meRst.Match()+nextRst.Match(), nextRst.Next(), true)
		}
		return NewRst(meRst.Match()+nextRst.Match(), nextRst.Next(), false)
	}
	return NewRst(meRst.Match(), meRst.Next(), false)
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

func (me *Com) NotsCheck(text string) bool {
	for _, not := range me.NextNots() {
		if rst := not.BeginningOfX(text); rst.SuccessFull() {
			return false
		}
	}
	return true
}
