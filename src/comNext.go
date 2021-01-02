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
