package main

type Other struct {
	Com
}

func (me *Other) BeginningOf(text string) *Rst {
	return NewRst("", text, true)
}

func NewOther() *Other {
	rst := &Other{}
	rst.Com = NewCom(rst, nil)
	return rst
}

func (me *Other) Print() {

}
