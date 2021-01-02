package main

type Def struct {
	Com
}

func (me *Def) BeginningOf(text string) *Rst {
	return NewRst("", text, true)
}

func NewDef() *Def {
	rst := &Def{}
	rst.Com = NewCom(rst, nil)
	return rst
}

func (me *Def) Print() {

}
