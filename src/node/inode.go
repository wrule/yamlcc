package node

// INode 接口
type INode interface {
	// 节点实例
	Me() INode
	// 获取上一个节点
	Prev() INode
	// 获取上n个节点
	PrevN(int) INode
	// 获取定义节点
	GetDef(string) *Def
	// 设置上一个节点
	SetPrev(INode)
	// 获取下节点列表
	Nexts() []INode
	// 获取下定义节点映射
	NextDefs() map[string]*Def
	// 获取下逻辑节点列表
	NextLogs() []INode
	// 获取下非命令节点列表
	NextNots() []*Not
	// 获取下其他命令节点
	NextOther() *Other
	// 设置下节点
	SetNexts([]INode)
	// 根据索引设置下节点
	SetNextByIndex(int, INode)
	// 追加下节点
	AppendNexts(INode)
	// 连接下节点（调用下节点追加）
	Link(INode)
	// 连接多个下节点
	Links([]INode)

	// 节点编译前的原始值
	SrcValue() interface{}
	// 是否是结束命令节点
	IsEnd() bool
	// 是否是非命令节点
	IsNot() bool
	// 是否是其他命令节点
	IsOther() bool
	// 是否是命令节点
	IsCmd() bool
	// 是否是定义节点
	IsDef() bool
	// 是否是逻辑节点
	IsLog() bool
	// 下节点列表是否为空
	IsNextsEmpty() bool
	// 下逻辑节点列表是否为空（相当于下逻辑节点只有一个结束命令节点）
	IsNextLogsEmpty() bool

	// 本节点头部匹配
	BeginningOf(string) *Rst
	// 本节点头部下推匹配（支持修整）
	BeginningOfX(string, bool) *Rst
	// 子节点头部下推匹配（支持修整）
	NextsBeginningOfX(string, bool) *Rst

	// 打印节点调试信息
	Print()
}
