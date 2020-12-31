package node

// BeginningOf s
func (me *Com) BeginningOf(text string) *Rst {
	return me.Me().BeginningOf(text)
}

// BeginningOfX s
func (me *Com) BeginningOfX(text string, trimHead bool) *Rst {
	return nil
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
