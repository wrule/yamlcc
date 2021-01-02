package main

type Back struct {
	Com
}

func (me *Back) BeginningOf(text string) *Rst {
	return NewRst("", text, true)
}

func NewBack() *Back {
	rst := &Back{}
	rst.Com = NewCom(rst, nil)
	return rst
}

func (me *Back) Print() {

}
