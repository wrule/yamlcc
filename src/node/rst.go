package node

// Rst 匹配结果
type Rst struct {
	success bool
	match   string
	next    string
}

// Success 匹配是否成功
func (me *Rst) Success() bool {
	return me.success
}

// Match 匹配获得文本
func (me *Rst) Match() string {
	return me.match
}

// Next 匹配剩余文本
func (me *Rst) Next() string {
	return me.next
}

// NewRst 构造函数
func NewRst(match, next string, success bool) *Rst {
	return &Rst{
		success: success,
		match:   match,
		next:    next,
	}
}
