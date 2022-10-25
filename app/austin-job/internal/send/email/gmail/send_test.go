package gmail

import (
	"austin-go/app/austin-job/internal/send/email"
	"testing"
)

var (
	cli *Client
)

func init() {
	config := &email.Config{
		Username:  "13834563@qq.com",
		Password:  "stqqvkpwrdegbgdj",
		Host:      "smtp.qq.com",
		Port:      465,
		Aliasname: "wander",
	}
	cli = NewClient(*config)
}

func TestClient_Send(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		msg := email.NewMessage("wander112900@qq.com", "hello", "hello")
		if err := cli.Send(msg, nil); err != nil {
			t.Errorf("Client.Send() error = %v, wantErr %v", err, "error")
		}
	})
}
