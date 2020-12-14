package node

// Dict 字典节点
type Dict struct {
	value   map[interface{}]interface{}
	nodeMap map[INode]INode
	prev    INode
	next    INode
	*Com
}

// Type 类型
func (me *Dict) Type() ENodeType {
	return NodeTypeDict
}

// BeginningOf s
func (me *Dict) BeginningOf(text string) (string, string) {
	return "", ""
}

// NewDict 构造函数
func NewDict(value map[interface{}]interface{}) *Dict {
	return &Dict{
		value: value,
	}
}
