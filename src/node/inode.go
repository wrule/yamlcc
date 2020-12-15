package node

// INode 节点接口
type INode interface {
	// 节点类型
	Type() ENodeType
	// 字符串匹配
	BeginningOf(string) (string, string)

	// 节点原始值
	SrcValue() interface{}
	// 上一个节点
	Prev() INode
	// 上n个节点
	PrevN(int) INode
	// 设置上一个节点
	SetPrev(INode)
	// 下一个节点
	Next() INode
	// 设置下一个节点
	SetNext(INode)
	// 根据名称获取节点定义
	GetDef(string) INode
}
