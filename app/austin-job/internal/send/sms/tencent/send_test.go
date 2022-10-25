package tencent

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
		TencentSMS Config
	}
	conf.LoadConfig("../../../../../../config.yaml", &c)
	cli = NewClient(c.TencentSMS)
}

func TestClient_Send(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		data := []string{
			"123456",
			"5",
		}
		msg := NewMessage([]string{"13388612969"}, "1580789", data, "1400753693", "涵睿科技")
		if err := cli.Send(msg, func(res interface{}) {
			fmt.Println(res)
		}); err != nil {
			t.Errorf("Client.Send() error = %v, wantErr %v", err, "error")
		}
	})
}
