package tencent

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Config struct {
	Url       string `json:"url"`
	Region    string `json:"region"` //地域
	SecretId  string `json:"secretId"`
	SecretKey string `json:"secretKey"`
}
