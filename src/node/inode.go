package node

// INode 节点接口
type INode interface {
	// 头部匹配（返回匹配到的字符串，剩余字符串，匹配结果）
	BeginningOf(string) (string, string, bool)

	// 头部匹配（忽略头部无效字符）
	BeginningTrimOf(string) (string, string, bool)
	// 原始值
	SrcValue() interface{}
	// 节点实例本身（用于在抽象类里访问具体类）
	Me() INode
	// IsEnd 判断节点是否是结束命令节点
	IsEnd() bool
	// IsEnd 判断节点是否是字典节点
	IsDict() bool
	// 上一个节点
	Prev() INode
	// 上n个节点
	PrevN(int) INode
	// 上n个字典节点
	PrevDictN(int) *Dict
	// 设置上一个节点
	SetPrev(INode)
	// 下一个节点
	Next() INode
	// 设置下一个节点
	SetNext(INode)
	// 根据名称获取定义的节点
	GetDef(string) INode
	// 根据名称获取定义的正则节点
	GetDefReg(string) *Reg
	// 打印节点信息
	Print()
}
