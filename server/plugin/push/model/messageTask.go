package model

// MessageTask 表示消息任务
type MessageTask struct {
	TemplateID     int      // 消息模板ID
	BusinessID     int      // 业务ID
	Recipients     []string // 接收者
	SenderIDType   string   // 发送ID类型
	Channel        string   // 发送渠道
	TemplateType   string   // 模板类型
	MessageType    string   // 消息类型
	BlockType      string   // 屏蔽类型
	Copywriting    string   // 文案类型
	SendingAccount string   // 发送账号
}
