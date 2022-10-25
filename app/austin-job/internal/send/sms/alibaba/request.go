package alibaba

// Message implements send.Message
type Request struct {
	//手机号码，多个用【,】分隔
	PhoneNumbers string
	//签名
	SignName string
	//模版ID
	TemplateCode string
	//模版参数json
	TemplateParam string
}

func NewMessage(PhoneNumbers, TemplateCode, TemplateParam, SignName string) *Request {
	return &Request{
		PhoneNumbers:  PhoneNumbers,
		TemplateCode:  TemplateCode,
		TemplateParam: TemplateParam,
		SignName:      SignName,
	}
}

func (m *Request) To() string {
	return m.PhoneNumbers
}

func (m *Request) Content() []byte {
	b, _ := json.Marshal(m)
	return b
}
