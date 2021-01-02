package main

type Root struct {
	Com
}

func (me *Root) BeginningOf(text string) *Rst {
	return NewRst("", text, true)
}

func NewRoot() *Root {
	rst := &Root{}
	rst.Com = NewCom(rst, nil)
	return rst
}

func (me *Root) Print() {

}
