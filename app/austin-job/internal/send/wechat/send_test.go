package wechat

import (
	"testing"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var (
	cli *Client
)

func init() {
	var c struct {
		Wechat Config
		Redis  redis.RedisConf
	}
	conf.LoadConfig("../../../../../config.yaml", &c)
	rcli := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})
	cli = NewClient(c.Wechat, *rcli)
}

func TestClient_Send(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		tlpdata := map[string]*TemplateDataItem{
			"first":    {Value: "这是一个测试订单", Color: "#173177"},
			"keyword1": {Value: "SL112313123123", Color: "#173177"},
			"remark":   {Value: "有新的订单来了，快发货吧。", Color: "#071D42"},
		}
		msg := NewRequest("oyggS5xOHgKvYo_f2GlZQexBOick", "j-OfIahoJGKC1hHCUEU-XapusCHzL6KTN9D3ntHgOD0", "https://www.qq.com", tlpdata)
		c := &Client{
			cfg:         cli.cfg,
			httpCli:     cli.httpCli,
			redisClient: cli.redisClient,
		}
		if err := c.Send(msg, nil); err != nil {
			t.Errorf("Client.Send() error = %v, wantErr %v", err, "error")
		}
	})
}
