package main

type Com struct {
	srcValue interface{}
	me       INode
	prev     INode
	nexts    []INode
}

func (me *Com) Me() INode {
	return me.me
}

func (me *Com) Prev() INode {
	return me.prev
}

func (me *Com) Nexts() []INode {
	return me.nexts
}

// NewCom 构造函数
func NewCom(me INode, srcValue interface{}) Com {
	return Com{
		me:       me,
		srcValue: srcValue,
	}
}
