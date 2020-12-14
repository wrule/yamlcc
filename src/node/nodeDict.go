package node

// Dict 字典节点
type Dict struct {
	value map[interface{}]interface{}
	nodes []INode
	prev  INode
	next  INode
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
	nodes := []INode{}
	for _key, _value := range value {
		node := BuildNode(_key)
		leafNode := BuildLeafNode(_value)
		Link(node, leafNode)
		nodes = append(nodes, node)
	}
	return &Dict{
		value: value,
		nodes: nodes,
	}
}
