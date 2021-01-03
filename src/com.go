package main

type Com struct {
	srcValue  interface{}
	me        INode
	prev      INode
	nexts     []INode
	nextDefs  map[string]*Def
	nextLogs  []INode
	nextOther INode
	nextNots  []INode
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

func (me *Com) Me() INode {
	return me.me
}

func (me *Com) Fasten(next INode) {
	me.AppendNexts(next)
	next.SetPrev(me.Me())
}

func (me *Com) Fastens(nexts []INode) {
	for _, next := range nexts {
		me.Fasten(next)
	}
}

func (me *Com) Link() {

}

// NewCom 构造函数
func NewCom(me INode, srcValue interface{}) Com {
	return Com{
		me:       me,
		srcValue: srcValue,
	}
}
