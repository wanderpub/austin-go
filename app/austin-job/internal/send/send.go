package send

type DoRes func(res interface{})

// Sender 用于发送信息
type Sender interface {
	// 发送信息到指定的客户端，to字段设置接收者地址，如有多个接收地址
	// 需要用';'隔开，msg为需要发送的信息
	Send(msg Message, do DoRes) error
}

// Message 发送的消息
type Message interface {
	// 发送的内容
	Content() []byte

	To() string
}
