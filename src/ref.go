package main

type Ref struct {
	Com
}

func (me *Ref) BeginningOf(text string) *Rst {
	return NewRst("", text, true)
}

func NewRef() *Ref {
	rst := &Ref{}
	rst.Com = NewCom(rst, nil)
	return rst
}

func (me *Ref) Print() {

}
