package node

import "fmt"

// Ref 引用节点
type Ref struct {
	defName string
	Com
}

// DefName 定义名称
func (me *Ref) DefName() string {
	return me.defName
}

// BeginningOf 匹配
func (me *Ref) BeginningOf(text string) *Rst {
	return me.GetDef(me.defName).BeginningOf(text)
}

// BeginningTrimOf s
func (me *Ref) BeginningTrimOf(text string) *Rst {
	return me.GetDef(me.defName).BeginningTrimOf(text)
}

// NewRef 构造函数
func NewRef(text string) *Ref {
	rst := &Ref{
		defName: text[1:],
	}
	rst.Com = NewCom(rst, text)
	return rst
}

// Print s
func (me *Ref) Print() {
	fmt.Printf("引用节点: %v\n", me.DefName())
}
