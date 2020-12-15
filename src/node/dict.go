package node

import "fmt"

// Dict 字典节点
type Dict struct {
	nodeMap    map[INode]INode
	defNodeMap map[string]INode
	logNodeMap map[INode]INode
	Com
}

// NodeMap 节点映射
func (me *Dict) NodeMap() map[INode]INode {
	return me.nodeMap
}

// DefNodeMap 定义节点映射
func (me *Dict) DefNodeMap() map[string]INode {
	return me.defNodeMap
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

func getDefNodeMap(nodeMap map[INode]INode) map[string]INode {
	rst := map[string]INode{}
	for key, value := range nodeMap {
		if key.Type() == NodeTypeDef {
			defNode := key.(*Def)
			rst[defNode.DefName()] = value
		}
	}
	return rst
}

// NewDict 构造函数
func NewDict(srcMap map[interface{}]interface{}) *Dict {
	dict := &Dict{
		Com: Com{srcValue: srcMap},
	}
	dict.nodeMap = getNodeMap(srcMap, dict)
	dict.defNodeMap = getDefNodeMap(dict.nodeMap)
	return dict
}

// Print 打印字典信息
func (me *Dict) Print() {
	fmt.Println("定义节点:")
	for key := range me.DefNodeMap() {
		fmt.Printf("\t%s\n", key)
	}
	fmt.Println("逻辑节点:")
	for key := range me.NodeMap() {
		fmt.Printf("\t%v\n", key)
	}
}
