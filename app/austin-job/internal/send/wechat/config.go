package wechat

import (
	jsoniter "github.com/json-iterator/go"
)

const (
	WeiXinAccessToken = "weixin_access_token"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Config struct {
	APPId     string
	APPSecret string
}
