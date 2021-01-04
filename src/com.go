package main

type Com struct {
	srcValue  interface{}
	me        INode
	prev      INode
	nexts     []INode
	nextDefs  map[string]*Def
	nextLogs  []INode
	nextOther *Other
	nextNots  []*Not
}

func (me *Com) SrcValue() interface{} {
	return me.srcValue
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

func (me *Com) Init() {
	me.updateNexts()
}

func (me *Com) Link() {

}

func (me *Com) GetDef(name string) *Def {
	curNode := me.Me()
	for curNode != nil {
		if def, found := curNode.NextDefs()[name]; found {
			return def
		}
		curNode = curNode.Prev()
	}
	panic("main.Com.GetDef: 找不到定义节点 " + name)
}

// NewCom 构造函数
func NewCom(me INode, srcValue interface{}) Com {
	return Com{
		me:       me,
		srcValue: srcValue,
	}
}
