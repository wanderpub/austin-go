package gmail

import (
	"austin-go/app/austin-job/internal/send"
	"austin-go/app/austin-job/internal/send/email"
	"fmt"
	"strings"

	"gopkg.in/gomail.v2"
)

type Client struct {
	cfg    email.Config
	dialer gomail.Dialer
}

func NewClient(cfg email.Config) *Client {
	if cfg.Aliasname == "" {
		cfg.Aliasname = "官方"
	}
	//创建SMTP客户端，连接到远程的邮件服务器，需要指定服务器地址、端口号、用户名、密码，如果端口号为465的话，
	d := gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password) // 设置邮件正文
	return &Client{
		cfg:    cfg,
		dialer: *d,
	}
}

// Send 发送邮件，msg用mail.NewRequest(...)生成
// do参数不做处理
func (c *Client) Send(msg send.Message, do send.DoRes) error {
	var r *email.Request
	var ok bool
	if r, ok = msg.(*email.Request); !ok {
		return fmt.Errorf("this type is not supported")
	}
	m := gomail.NewMessage(
		//就选择utf-8格式保存
		gomail.SetEncoding(gomail.Base64),
	)
	mailTo := strings.Split(r.To(), ";")
	m.SetHeader("From", m.FormatAddress(c.cfg.Username, c.cfg.Aliasname)) // 添加别名
	m.SetHeader("To", mailTo...)                                          // 发送给用户(可以多个)
	m.SetHeader("Subject", r.Subject())                                   // 设置邮件主题
	contentType := r.TextType(r.Data())
	m.SetBody(contentType, r.Data())
	/*
		//一个文件（加入发送一个 txt 文件）：/tmp/foo.txt，我需要将这个文件以邮件附件的方式进行发送，同时指定附件名为：附件.txt
		//同时解决了文件名乱码问题
		name := "附件.txt"
		m.Attach("E:/GoCode/src/goMail/gomail.txt",
			gomail.Rename(name), //重命名
			gomail.SetHeader(map[string][]string{
				"Content-Disposition": []string{
					fmt.Sprintf(`attachment; filename="%s"`, mime.QEncoding.Encode("UTF-8", name)),
				},
			}),
		)
	*/

	err := c.dialer.DialAndSend(m)
	if err != nil {
		return err
	} else {
		if do != nil {
			do(true)
		}
	}
	return err
}
