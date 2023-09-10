package setting

import (
	"time"
)

 // ServerSettingS is a struct that holds the server settings
type ServerSettingS struct {
	// RunMode is a string that holds the run mode of the server
	RunMode           string
	// HttpPort is a string that holds the http port of the server
	HttpPort          string
	// ReadTimeout is a duration that holds the read timeout of the server
	ReadTimeout       time.Duration
	// WriteTimeout is a duration that holds the write timeout of the server
	WriteTimeout      time.Duration
	// OpenServiceListen is a string that holds the open service listen address of the server
	OpenServiceListen string
	// PrivateHttpListen is a string that holds the private http listen address of the server
	PrivateHttpListen string
	// DefaultAccessHost is a string that holds the default access host of the server
	DefaultAccessHost string
	// SignExpiry is an int64 that holds the sign expiry of the server
	SignExpiry        int64
}

type AppSettingS struct {
	// 默认页面大小
	DefaultPageSize       int
	// 最大页面大小
	MaxPageSize           int
	// 默认上下文超时时间
	DefaultContextTimeout time.Duration
	// 日志保存路径
	LogSavePath           string
	// 日志文件名
	LogFileName           string
	// 日志文件扩展名
	LogFileExt            string
	// 上传保存路径
	UploadSavePath        string
	// 上传服务器URL
	UploadServerUrl       string
	// 上传图片最大尺寸
	UploadImageMaxSize    int
	// 上传图片允许的扩展名
	UploadImageAllowExts  []string
}

type EmailSettingS struct {
	// Host is the hostname of the email server
	Host     string
	// Port is the port number of the email server
	Port     int
	// UserName is the username for the email server
	UserName string
	// Password is the password for the email server
	Password string
	// IsSSL is a boolean value indicating whether the email server is SSL or not
	IsSSL    bool
	// From is the email address that will be used as the sender of the email
	From     string
	// To is an array of email addresses that will be used as the recipients of the email
	To       []string
}

type JWTSettingS struct {
	// Secret is the secret used to sign the JWT
	Secret string
	// Issuer is the issuer of the JWT
	Issuer string
	// Expire is the duration of the JWT
	Expire time.Duration
}
// SecuritySettingS struct
type SecuritySettingS struct {
	// TokenAuthKey 认证密钥
	TokenAuthKey      string
	// EnableHtmlEncrypt 是否启用html加密
	EnableHtmlEncrypt bool
	// HtmlEncryptKey html加密密钥
	HtmlEncryptKey    string
}

type DatabaseSettingS struct {
	// 数据库类型
	DBType       string
	// 用户名
	UserName     string
	// 密码
	Password     string
	// 主机
	Host         string
	// 数据库名
	DBName       string
	// 表前缀
	TablePrefix  string
	// 字符集
	Charset      string
	// 是否解析时间
	ParseTime    bool
	// 最大闲置连接数
	MaxIdleConns int
	// 最大打开连接数
	MaxOpenConns int
}

 // This struct contains the host, password, and db information for a Redis instance
type RedisS struct {
	Host     string
	Password string
	Db       int
}
type RabbitMQS struct {
	// Dsn is the RabbitMQ connection string
	Dsn string
}
type WechatPaySettingS struct {
	// 应用ID
	Appid                       string
	// 商户号
	Mchid                       string
	// 商户证书序列号
	MchCertificateSerial        string
	// 商户私钥文件路径
	MchPrivateKeyFilePath       string
	// 平台证书文件路径
	PlatformCertificateFilePath string
	// apiv3密钥
	Apiv3Key                    string
}
type AlipaySettingS struct {
	// Appid
	Appid        string
	// AliPublicKey
	AliPublicKey string
	// PrivateKey
	PrivateKey   string
	// IsProd
	IsProd       bool
	// SecretKey
	SecretKey    string
}

// OSSSettingS struct
type OSSSettingS struct {
	// BucketName is the name of the bucket
	BucketName      string
	// AccessURLDomain is the access URL domain
	AccessURLDomain string
	// Endpoint is the endpoint of the OSS
	Endpoint        string
	// AccessKeyID is the access key ID
	AccessKeyID     string
	// AccessKeySecret is the access key secret
	AccessKeySecret string
}
var sections = make(map[string]interface{})

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}
