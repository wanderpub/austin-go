package email

import (
	"testing"

	"github.com/zeromicro/go-zero/core/conf"
)

var (
	cli *Client
)

func init() {
	var c struct {
		Email Config
	}
	conf.LoadConfig("../../../../../config.yaml", &c)
	cli = NewClient(c.Email)
}

func TestClient_Send(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		msg := NewMessage("wander112900@qq.com", "hello", "hello")
		if err := cli.Send(msg, nil); err != nil {
			t.Errorf("Client.Send() error = %v, wantErr %v", err, "error")
		}
	})
}
