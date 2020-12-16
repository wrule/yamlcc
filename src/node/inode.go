package node

// INode 节点接口
type INode interface {
	// 字符串匹配
	BeginningOf(string) (string, string, bool)

	// IsEnd 判断节点是否是结束命令节点
	IsEnd() bool
	// 字符串匹配（忽略无效字符）
	BeginningTrimOf(string) (string, string, bool)
	// 原始值
	SrcValue() interface{}
	// 节点实例本身
	Me() INode
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
	// 根据名称获取定义节点
	GetDef(string) INode
	// 根据名称获取正则节点
	GetDefReg(string) *Reg
	// 打印节点信息
	Print()
}
