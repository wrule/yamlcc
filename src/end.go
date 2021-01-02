package main

type End struct {
	Com
}

func (me *End) BeginningOf(text string) *Rst {
	return NewRst("", text, true)
}

func NewEnd() *End {
	rst := &End{}
	rst.Com = NewCom(rst, nil)
	return rst
}

func (me *End) Print() {

}
