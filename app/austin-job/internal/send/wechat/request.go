package wechat

// TemplateMessage 发送的模板消息内容
type Request struct {
	ToUser     string                       `json:"touser"`          // 必须, 接受者OpenID
	TemplateId string                       `json:"template_id"`     // 必须, 模版ID
	URL        string                       `json:"url,omitempty"`   // 可选, 用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中
	Color      string                       `json:"color,omitempty"` // 可选, 整个消息的颜色, 可以不设置
	Data       map[string]*TemplateDataItem `json:"data"`            // 必须, 模板数据

	MiniProgram struct {
		AppID    string `json:"appid"`    // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系）
		PagePath string `json:"pagepath"` // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar）
	} `json:"miniprogram"` // 可选,跳转至小程序地址
}

// TemplateDataItem 模版内某个 .DATA 的值
type TemplateDataItem struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

func NewRequest(to, templateId, url string, data map[string]*TemplateDataItem) *Request {
	return &Request{
		ToUser:     to,
		TemplateId: templateId,
		URL:        url,
		Data:       data,
	}
}

func (r *Request) Content() []byte {
	b, _ := json.Marshal(r)
	return b
}

func (r *Request) To() string {
	return r.ToUser
}
