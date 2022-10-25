package wechat

import (
	"austin-go/app/austin-job/internal/send"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	defaultTimeout = time.Second * 10
	lockKey        = WeiXinAccessToken + "_lock"
	tokenURL       = "https://api.weixin.qq.com/cgi-bin/token"
	sendURL        = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="
)

var (
	ErrTokenOverdue = errors.New("weixin token is overdued")
)

type Client struct {
	httpCli     *http.Client
	cfg         Config
	redisClient *redis.Redis
}

func NewClient(cfg Config, redis redis.Redis) *Client {
	return &Client{
		cfg:         cfg,
		redisClient: &redis,
		httpCli: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

// accessTokenData 先从缓存中获取token，如果不存在再http请求获取
func (c *Client) accessTokenData() (string, error) {
	var count int
TOKEN:
	token, err := c.token()
	if err == nil {
		return token, nil
	}
	// 获取分布式锁
	_, err = c.lockTokenGet()
	if err != nil {
		fmt.Println("获取weixin-token的分布式锁失败")
		// 如果超过三次还没有从缓存中获取到token，那么就直接去请求token
		if count > 2 {
			fmt.Println("获取weixin-token分布式锁失败次数过多，直接去请求连接获取")
			goto GET_TOKEN
		}
		// 如果已有线程在更新token，那么其他线程会进入等待，等待后会重新去缓存中获取token
		// 防止竞争，随机设置等待时间
		sleepTime := 1000 + rand.Intn(1000)
		time.Sleep(time.Millisecond * time.Duration(sleepTime))
		count++
		goto TOKEN
	}
GET_TOKEN:
	result, err := c.requestAccessToken()
	if err != nil {
		return "", err
	}
	c.storeToken(result.AccessToken, result.ExpiresIn-100)
	c.unLockTokenGet()
	return result.AccessToken, nil
}

// requestAccessToken 通过http请求获取weixin_token
func (c *Client) requestAccessToken() (*Response, error) {
	url := fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", tokenURL, c.cfg.APPId, c.cfg.APPSecret)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := &Response{}
	err = json.Unmarshal(b, result)
	if err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, fmt.Errorf("get accessTokenData failed,code: %d error: %v", result.ErrCode, result.ErrMsg)
	}
	return result, nil
}

// token 从缓存中获取token
func (c *Client) token() (string, error) {
	b, err := c.redisClient.GetCtx(context.Background(), WeiXinAccessToken)
	if err != nil {
		return "", err
	}
	if b == "" {
		return "", fmt.Errorf("未获取到TOKEN")
	}
	fmt.Println(b)
	return b, nil
}

// storeToken 在缓存中存储token
func (c *Client) storeToken(v string, e int) error {
	return c.redisClient.SetexCtx(context.Background(), WeiXinAccessToken, v, e)
}

func (c *Client) removeToken() {
	c.redisClient.DelCtx(context.Background(), WeiXinAccessToken)
}

// 分布式锁，防止多进程或协程去同时去获取accesstoken
func (c *Client) lockTokenGet() (bool, error) {
	return c.redisClient.SetnxExCtx(context.Background(), lockKey, "lock", 10)
}

func (c *Client) unLockTokenGet() {
	c.redisClient.DelCtx(context.Background(), lockKey)
}

func (c *Client) Send(msg send.Message, do send.DoRes) error {
	token, err := c.accessTokenData()
	if err != nil {
		return err
	}
	url := sendURL + token
	// 微信token失效返回json {"errcode":40001,"errmsg":"invalid credential, access_token is invalid or not latest hint: [CVJXrA0584vr61!]"}
	r := bytes.NewReader(msg.Content())
	req, err := http.NewRequest("POST", url, r)
	if err != nil {
		return err
	}
	resp, err := c.httpCli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	result := &Response{}
	err = json.Unmarshal(b, result)
	if err != nil {
		return err
	}
	if result.ErrCode == 40001 {
		// token超时，移除token
		c.removeToken()
		return ErrTokenOverdue
	}
	if result.ErrCode != 0 {
		return fmt.Errorf("send msg failed,code: %d error: %v", result.ErrCode, result.ErrMsg)
	}
	if do != nil {
		do(result)
	}
	return nil
}
