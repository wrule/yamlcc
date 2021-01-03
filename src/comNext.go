package main

func (me *Com) Nexts() []INode {
	return me.nexts
}

func (me *Com) AppendNexts(next INode) {
	me.nexts = append(me.nexts, next)
}

func (me *Com) NextsIsEmpty() bool {
	return len(me.Nexts()) < 1
}

func (me *Com) updateNextDefs() {
	me.nextDefs = map[string]*Def{}
	for _, node := range me.nexts {
		if node.IsDef() {
			def := node.(*Def)
			me.nextDefs[def.DefName()] = def
		}
	}
}

func (me *Com) updateNextLogs() {
	me.nextLogs = []INode{}
	for _, node := range me.nexts {
		if node.IsLog() {
			me.nextLogs = append(me.nextLogs, node)
		}
	}
}

func (me *Com) updateNextOther() {
	for _, node := range me.nexts {
		if node.IsOther() {
			other := node.(*Other)
			me.nextOther = other
			// other只采集收集到的第一个
			return
		}
	}
}

func (me *Com) updateNextNots() {
	me.nextNots = []*Not{}
	for _, node := range me.nexts {
		if node.IsNot() {
			not := node.(*Not)
			me.nextNots = append(me.nextNots, not)
		}
	}
}

func (me *Com) updateNexts() {
	me.updateNextDefs()
	me.updateNextLogs()
	me.updateNextOther()
	me.updateNextNots()
}
