package email

// Config 用于连接邮件服务器的配置
type Config struct {
	// 邮件服务的域名
	Host string
	//端口
	Port int
	// 登陆邮件服务的用户名
	Username string
	//别名
	Aliasname string
	// 登陆密码
	Password string
	//是否tls
	TLS bool
	//邮件服务的域名:端口(选填)
	ServerAddr string
}
