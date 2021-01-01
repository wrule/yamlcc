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
	return nil
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
