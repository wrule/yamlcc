package main

type Node struct {
	srcMap map[string]interface{}
}

// NewNode 构造函数
func NewNode(srcMap map[string]interface{}) *Node {
	return &Node{
		srcMap: srcMap,
	}
}
