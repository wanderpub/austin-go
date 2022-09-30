package config

import (
	"austin-go/common/dbx"
	"austin-go/common/mq"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql      dbx.DbConfig
	Rabbit     mq.RabbitConf
	CacheRedis cache.CacheConf
}
