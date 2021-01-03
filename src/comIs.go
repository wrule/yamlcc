package main

// IsRoot 是否是根节点
func (me *Com) IsRoot() bool {
	_, ok := me.Me().(*Root)
	return ok
}

// IsReg 是否是正则节点
func (me *Com) IsReg() bool {
	_, ok := me.Me().(*Reg)
	return ok
}

// IsDef 是否是定义节点
func (me *Com) IsDef() bool {
	_, ok := me.Me().(*Def)
	return ok
}

// IsRef 是否是引用节点
func (me *Com) IsRef() bool {
	_, ok := me.Me().(*Ref)
	return ok
}

// IsBack 是否是回跳节点
func (me *Com) IsBack() bool {
	_, ok := me.Me().(*Back)
	return ok
}

// IsEnd 是否是结束命令节点
func (me *Com) IsEnd() bool {
	_, ok := me.Me().(*End)
	return ok
}

// IsOther 是否是其他命令节点
func (me *Com) IsOther() bool {
	_, ok := me.Me().(*Other)
	return ok
}

// IsNot 是否是非命令节点
func (me *Com) IsNot() bool {
	_, ok := me.Me().(*Not)
	return ok
}
