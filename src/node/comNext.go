package node

// SetNexts s
func (me *Com) SetNexts(nexts []INode) {
	me.nexts = nexts
	me.updateNexts()
}

// AppendNexts s
func (me *Com) AppendNexts(next INode) {
	me.nexts = append(me.nexts, next)
	me.updateNexts()
}

// Nexts s
func (me *Com) Nexts() []INode {
	return me.nextLogs
}

// NextDefs s
func (me *Com) NextDefs() map[string]*Def {
	return me.nextDefs
}

// NextCmds s
func (me *Com) NextCmds() []INode {
	return me.nextCmds
}

// NextNots s
func (me *Com) NextNots() []*Not {
	return me.nextNots
}

// NextOther s
func (me *Com) NextOther() *Other {
	return me.nextOther
}

// updateNextDefs 同步更新下节点的定义节点列表
func (me *Com) updateNextDefs() {
	me.nextDefs = map[string]*Def{}
	for _, node := range me.nexts {
		if node.IsDef() {
			def := node.(*Def)
			me.nextDefs[def.DefName()] = def
		}
	}
}

// updateNextLogs 同步更新下节点的逻辑节点列表
func (me *Com) updateNextLogs() {
	me.nextLogs = []INode{}
	for _, node := range me.nexts {
		if node.IsLog() {
			me.nextLogs = append(me.nextLogs, node)
		}
	}
}

// updateNextCmds 同步更新下节点的命令节点列表
func (me *Com) updateNextCmds() {
	me.nextCmds = []INode{}
	for _, node := range me.nexts {
		if node.IsCmd() {
			me.nextCmds = append(me.nextCmds, node)
		}
	}
}

// updateNextNots 同步更新下节点的Not节点列表
func (me *Com) updateNextNots() {
	me.nextNots = []*Not{}
	for _, node := range me.nexts {
		if node.IsNot() {
			not := node.(*Not)
			me.nextNots = append(me.nextNots, not)
		}
	}
}

// updateNextOther 同步更新下节点的Other节点
func (me *Com) updateNextOther() {
	for _, node := range me.nexts {
		if node.IsOther() {
			other := node.(*Other)
			me.nextOther = other
			return
		}
	}
}

// updateNexts 同步更新其他的nexts相关数据结构
func (me *Com) updateNexts() {
	me.updateNextDefs()
	me.updateNextLogs()
	me.updateNextCmds()
	me.updateNextNots()
	me.updateNextOther()
}
