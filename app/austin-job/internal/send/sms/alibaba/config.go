package alibaba

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Config struct {
	AccessKeyId  string
	AccessSecret string
	GatewayURL   string
}
