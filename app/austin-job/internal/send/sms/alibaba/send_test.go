package alibaba

import (
	"fmt"
	"testing"

	"github.com/zeromicro/go-zero/core/conf"
)

var (
	cli *Client
)

func init() {
	var c struct {
		AliyunSMS Config
	}
	conf.LoadConfig("../../../../../../config.yaml", &c)
	fmt.Println(c)
	cli = NewClient(c.AliyunSMS)
}

func TestClient_Send(t *testing.T) {
	fmt.Println("开始测试发送代码")
	t.Run("test", func(t *testing.T) {
		/*
			msg := NewMessage("18667143169", "SMS_137955257", "{\"code\":1222}", "xxApp")
			if err := cli.Send(msg, func(res interface{}) {
				fmt.Println(res)
			}); err != nil {
				t.Errorf("Client.Send() error = %v, wantErr %v", err, "error")
			}
		*/
	})
}

func TestClient_GetTemplateList(t *testing.T) {
	fmt.Println("开始测试发送代码")
	res, err := cli.getTemplateList(1, 20)
	if err != nil {
		t.Errorf("Client.Send() error = %v", err)
	}
	fmt.Println(res)
}
