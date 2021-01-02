package main

func (me *Com) AppendNexts(next INode) {
	me.nexts = append(me.nexts, next)
}
