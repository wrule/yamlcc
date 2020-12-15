package node

// Dict 字典节点
type Dict struct {
	value map[interface{}]interface{}
	nodes []INode
	Com
}

// Nodes 获取子节点列表
func (me *Dict) Nodes() []INode {
	return me.nodes
}

// Type 类型
func (me *Dict) Type() ENodeType {
	return NodeTypeDict
}

// BeginningOf s
func (me *Dict) BeginningOf(text string) (string, string) {
	return "", ""
}

// getNodes 获取节点
func getNodes(
	srcMap map[interface{}]interface{},
) []INode {
	rst := []INode{}
	for key, value := range srcMap {
		node := BuildNode(key)
		leafNode := BuildLeafNode(value)
		Link(node, leafNode)
		rst = append(rst, node)
	}
	return rst
}

// NewDict 构造函数
func NewDict(value map[interface{}]interface{}) *Dict {
	dict := &Dict{
		value: value,
	}
	nodes := getNodes(dict.value)
	for _, node := range nodes {
		node.SetPrev(dict)
	}
	dict.nodes = nodes
	return dict
}
