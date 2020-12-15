package node

// Dict 字典节点
type Dict struct {
	srcMap  map[interface{}]interface{}
	nodeMap map[INode]INode
	Com
}

// Type 类型
func (me *Dict) Type() ENodeType {
	return NodeTypeDict
}

// BeginningOf 匹配
func (me *Dict) BeginningOf(text string) (string, string, bool) {
	return "", text, false
}

// getNodeMap 获取节点Map
func getNodeMap(
	srcMap map[interface{}]interface{},
	prev INode,
) map[INode]INode {
	rst := map[INode]INode{}
	for key, value := range srcMap {
		node := BuildNode(key)
		node.SetPrev(prev)
		leafNode := BuildLeafNode(value)
		Link(node, leafNode)
		rst[node] = leafNode
	}
	return rst
}

// NewDict 构造函数
func NewDict(srcMap map[interface{}]interface{}) *Dict {
	dict := &Dict{
		srcMap: srcMap,
	}
	dict.nodeMap = getNodeMap(dict.srcMap, dict)
	return dict
}
