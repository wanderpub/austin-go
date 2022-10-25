package email

import (
	"austin-go/app/austin-job/internal/send"
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"
)

type Client struct {
	cfg  Config
	auth smtp.Auth
}

func NewClient(cfg Config) *Client {
	cfg.ServerAddr = fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	if cfg.Port == 465 {
		cfg.TLS = true
	}
	if cfg.Aliasname == "" {
		cfg.Aliasname = "官方"
	}
	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)
	return &Client{
		cfg:  cfg,
		auth: auth,
	}
}

// Send 发送邮件，msg用mail.NewRequest(...)生成
// do参数不做处理
func (c *Client) Send(msg send.Message, do send.DoRes) error {
	var m *Request
	var ok bool
	if m, ok = msg.(*Request); !ok {
		return fmt.Errorf("this type is not supported, use sms.NewRequest()")
	}
	// 拼接发送的信息
	msgStr := fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type:%s;charset=UTF-8\r\n\r\n%s",
		c.cfg.Aliasname+" <"+c.cfg.Username+">",
		m.To(),
		m.Subject(),
		m.TextType(m.Data()),
		m.Data(),
	)
	if c.cfg.TLS {
		return c.sendTls(strings.Split(msg.To(), ";"), []byte(msgStr))
	}

	return smtp.SendMail(c.cfg.ServerAddr, c.auth, c.cfg.Username, strings.Split(msg.To(), ";"), []byte(msgStr))
}

func (c *Client) sendTls(tos []string, msg []byte) error {
	conn, err := tls.Dial("tcp", c.cfg.ServerAddr, nil)
	if err != nil {
		return err
	}
	cli, err := smtp.NewClient(conn, c.cfg.Host)
	if err != nil {
		return err
	}
	defer cli.Close()
	if ok, _ := cli.Extension("AUTH"); ok {
		if err := cli.Auth(c.auth); err != nil {
			return err
		}
	}
	if err = cli.Mail(c.cfg.Username); err != nil {
		return err
	}
	for _, to := range tos {
		if err = cli.Rcpt(to); err != nil {
			return err
		}
	}

	w, err := cli.Data()
	if err != nil {
		return err
	}
	if _, err = w.Write(msg); err != nil {
		return err
	}
	if err = w.Close(); err != nil {
		return err
	}
	return cli.Quit()
}
