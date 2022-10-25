package tencent

import (
	"austin-go/app/austin-job/internal/send"
	"fmt"
	"strings"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

const (
// endpoint = "sms.tencentcloudapi.com"
)

// Client 发送短信客户端，implements send.Sender
type Client struct {
	// cfg    Config
	client *sms.Client
}

func NewClient(cfg Config) *Client {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	// 密钥可前往https://console.cloud.tencent.com/cam/capi网站进行获取
	credential := common.NewCredential(
		cfg.SecretId,
		cfg.SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = cfg.Url //接口地址
	// 实例化要请求产品的client对象,clientProfile是可选的
	cli, _ := sms.NewClient(credential, cfg.Region, cpf)
	return &Client{
		// cfg:    cfg,
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
	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := sms.NewSendSmsRequest()
	mobiles := strings.Split(m.PhoneNumberSet, ";")
	request.PhoneNumberSet = common.StringPtrs(mobiles)
	request.SmsSdkAppId = common.StringPtr(m.SmsSdkAppId)
	request.SignName = common.StringPtr(m.SignName)
	request.TemplateId = common.StringPtr(m.TemplateId)
	request.TemplateParamSet = common.StringPtrs(m.TemplateParamSet)

	// 返回的resp是一个SendSmsResponse的实例，与请求对象对应
	response, err := c.client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return fmt.Errorf("API error %v, wantErr %v", response, "error")
	}
	if err != nil {
		return fmt.Errorf("Client.Send() error = %v, wantErr %v", err, "error")
	}
	if do != nil {
		res, _ := json.MarshalToString(response)
		do(res)
	}
	return nil
}
