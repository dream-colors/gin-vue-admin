package model

// Channel 表示通道模板
type Channel struct {
	ID         int    // 通道ID
	Name       string // 通道名称
	Type       string // 通道类型
	Parameters string // 通道参数
	Domain     string // 域名
	Enabled    bool   // 启用状态
}
