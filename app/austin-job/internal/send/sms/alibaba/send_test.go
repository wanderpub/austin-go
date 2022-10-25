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
		msg := NewRequest("18667143169", "锣号App", "CheckCode", "{\"code\":12}", "123")
		err := cli.Send(msg, func(res interface{}) {
			fmt.Println(res)
		})

		if err != nil {
			t.Errorf("Client.Send() error = %v", err)
		}
	})

	/*

		type fields struct {
			cfg    Config
			client *http.Client
		}
		type args struct {
			msg send.Message
			do  send.DoRes
		}
		tests := []struct {
			name    string
			fields  fields
			args    args
			wantErr bool
		}{
			// TODO: Add test cases.
			{
				name: "send_case_1",
				fields: fields{
					cfg:    cli.cfg,
					client: cli.client,
				},
				args: args{
					msg: NewRequest("18667143169", "test", "te", "te", "123"),
				},
				wantErr: true,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := &Client{
					cfg:    tt.fields.cfg,
					client: tt.fields.client,
				}
				if err := c.Send(tt.args.msg, tt.args.do); (err != nil) != tt.wantErr {
					t.Errorf("Client.Send() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	*/
}
