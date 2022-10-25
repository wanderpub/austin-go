package tencent

import "strings"

// Message implements send.Message
type Request struct {
	// 下发手机号码，采用 E.164 标准，格式为+[国家或地区码][手机号]，单次请求最多支持200个手机号且要求全为境内手机号或全为境外手机号。
	// 例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号。
	// 注：发送国内短信格式还支持0086、86或无任何国家或地区码的11位手机号码，前缀默认为+86。
	PhoneNumberSet string `json:"PhoneNumberSet,omitempty" name:"PhoneNumberSet"`

	// 短信 SdkAppId，在 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage)  添加应用后生成的实际 SdkAppId，示例如1400006666。
	SmsSdkAppId string `json:"SmsSdkAppId,omitempty" name:"SmsSdkAppId"`

	// 模板 ID，必须填写已审核通过的模板 ID。模板 ID 可前往 [国内短信](https://console.cloud.tencent.com/smsv2/csms-template) 或 [国际/港澳台短信](https://console.cloud.tencent.com/smsv2/isms-template) 的正文模板管理查看，若向境外手机号发送短信，仅支持使用国际/港澳台短信模板。
	TemplateId string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 短信签名内容，使用 UTF-8 编码，必须填写已审核通过的签名，例如：腾讯云，签名信息可前往 [国内短信](https://console.cloud.tencent.com/smsv2/csms-sign) 或 [国际/港澳台短信](https://console.cloud.tencent.com/smsv2/isms-sign) 的签名管理查看。
	// <dx-alert infotype="notice" title="注意">发送国内短信该参数必填。</dx-alert>
	SignName string `json:"SignName,omitempty" name:"SignName"`

	// 模板参数，若无模板参数，则设置为空。
	// <dx-alert infotype="notice" title="注意">模板参数的个数需要与 TemplateId 对应模板的变量个数保持一致。</dx-alert>
	TemplateParamSet []string `json:"TemplateParamSet,omitempty" name:"TemplateParamSet"`

	// 短信码号扩展号，默认未开通，如需开通请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	ExtendCode string `json:"ExtendCode,omitempty" name:"ExtendCode"`

	// 用户的 session 内容，可以携带用户侧 ID 等上下文信息，server 会原样返回。
	SessionContext string `json:"SessionContext,omitempty" name:"SessionContext"`

	// 国内短信无需填写该项；国际/港澳台短信已申请独立 SenderId 需要填写该字段，默认使用公共 SenderId，无需填写该字段。
	// 注：月度使用量达到指定量级可申请独立 SenderId 使用，详情请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。
	SenderId string `json:"SenderId,omitempty" name:"SenderId"`
}

func NewMessage(mobiles []string, templateId string, templateParam []string, appid string, signName string) *Request {
	mobile := strings.Join(mobiles, ";")
	return &Request{
		PhoneNumberSet:   mobile,
		TemplateId:       templateId,
		SmsSdkAppId:      appid,
		SignName:         signName,
		TemplateParamSet: templateParam,
	}
}

func (m *Request) To() string {
	return m.PhoneNumberSet
}

func (m *Request) TemplateID() string {
	return m.TemplateId
}

func (m *Request) SmsSdkAppID() string {
	return m.SmsSdkAppId
}

//SignName
func (m *Request) Sign() string {
	return m.SignName
}

//TemplateParamSet
func (m *Request) TemplateParam() []string {
	return m.TemplateParamSet
}

func (m *Request) Content() []byte {
	b, _ := json.Marshal(m)
	return b
}
