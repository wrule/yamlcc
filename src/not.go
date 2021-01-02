package main

type Not struct {
	Com
}

func (me *Not) BeginningOf(text string) *Rst {
	return NewRst("", text, true)
}

func NewNot() *Not {
	rst := &Not{}
	rst.Com = NewCom(rst, nil)
	return rst
}

func (me *Not) Print() {

}
