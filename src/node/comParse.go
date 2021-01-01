package node

// BeginningOf s
func (me *Com) BeginningOf(text string) *Rst {
	return me.Me().BeginningOf(text)
}

// BeginningOfX s
func (me *Com) BeginningOfX(text string, trimHead bool) *Rst {
	ivdRst := NewRst("", text, true)
	// 按需进行头部修整匹配
	if trimHead {
		ivdRst = me.Me().GetDef("invalid").BeginningOfX(text, false)
	}
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
	successList := []*Rst{}
	failureList := []*Rst{}

	// 非逻辑检查
	if me.NotsCheck(text, trimHead) == false {
		ivdRst := me.invalidMatch(text, trimHead)
		return NewRst(ivdRst.Match(), ivdRst.Next(), false)
	}

	// 遍历下推匹配逻辑节点，并且采集成功失败结果
	for _, log := range me.Me().NextLogs() {
		rst := log.BeginningOfX(text, trimHead)
		if rst.Success() {
			successList = append(successList, rst)
		} else {
			failureList = append(failureList, rst)
		}
	}
	return nil
}

func (me *Com) invalidMatch(text string, trimHead bool) *Rst {
	ivdRst := NewRst("", text, true)
	if trimHead {
		ivdRst = me.Me().GetDef("invalid").BeginningOfX(text, false)
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
