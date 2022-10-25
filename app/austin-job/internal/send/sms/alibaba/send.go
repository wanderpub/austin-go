package alibaba

import (
	"austin-go/app/austin-job/internal/send"
	"fmt"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	smsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
)

type Client struct {
	client *smsapi.Client
}

func NewClient(cfg Config) *Client {
	config := &openapi.Config{
		AccessKeyId:     &cfg.AccessKeyId,
		AccessKeySecret: &cfg.AccessSecret,
	}
	// 访问的域名
	config.Endpoint = tea.String(cfg.GatewayURL)
	cli, _ := smsapi.NewClient(config)
	return &Client{
		client: cli,
	}
}

// Send 发送短信，msg需要使用sms.NewRequest(...)生成
// 可以使用do将请求结果回传,也可以传nil忽略结果
func (c *Client) Send(msg send.Message, do send.DoRes) error {
	var m *Request
	var ok bool
	if m, ok = msg.(*Request); !ok {
		return fmt.Errorf("this type is not supported, use sms.NewRequest()")
	}
	request := &smsapi.SendSmsRequest{}
	request.SetPhoneNumbers(m.PhoneNumbers)
	request.SetSignName(m.SignName)
	request.SetTemplateCode(m.TemplateCode)
	request.SetTemplateParam(m.TemplateParam)
	response, err := c.client.SendSms(request)
	if err != nil {
		return fmt.Errorf("Client.Send() error = %v", err)
	}
	if do != nil {
		do(response)
	}
	return nil
}

func (c *Client) getTemplateList(pageIndex, pageSize int) (string, error) {
	request := &smsapi.QuerySmsTemplateListRequest{
		PageIndex: tea.ToInt32(&pageIndex),
		PageSize:  tea.ToInt32(&pageSize),
	}
	var res *smsapi.QuerySmsTemplateListResponse
	res, err := c.client.QuerySmsTemplateList(request)
	if err != nil {
		return res.Body.String(), fmt.Errorf("QuerySmsTemplateList error = %v", err)
	}

	return res.Body.String(), err
}
