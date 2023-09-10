package global

import (
	"github.com/haierspi/golang-image-upload-service/pkg/logger"
	"github.com/haierspi/golang-image-upload-service/pkg/setting"
)

var (
	// ServerSetting is a pointer to a setting.ServerSettingS
	ServerSetting         *setting.ServerSettingS
	// AppSetting is a pointer to a setting.AppSettingS
	AppSetting            *setting.AppSettingS
	// EmailSetting is a pointer to a setting.EmailSettingS
	EmailSetting          *setting.EmailSettingS
	// JWTSetting is a pointer to a setting.JWTSettingS
	JWTSetting            *setting.JWTSettingS
	// SecuritySetting is a pointer to a setting.SecuritySettingS
	SecuritySetting       *setting.SecuritySettingS
	// DatabaseSetting is a pointer to a setting.DatabaseSettingS
	DatabaseSetting       *setting.DatabaseSettingS
	// RedisSetting is a pointer to a setting.RedisS
	RedisSetting          *setting.RedisS
	// RabbitMQSetting is a pointer to a setting.RabbitMQS
	RabbitMQSetting       *setting.RabbitMQS
	// WechatJSAPIPaySetting is a pointer to a setting.WechatPaySettingS
	WechatJSAPIPaySetting *setting.WechatPaySettingS
	// WechatH5PaySetting is a pointer to a setting.WechatPaySettingS
	WechatH5PaySetting    *setting.WechatPaySettingS
	// AlipayH5PaySetting is a pointer to a setting.AlipaySettingS
	AlipayH5PaySetting    *setting.AlipaySettingS
	// Logger is a pointer to a logger.Logger
	Logger                *logger.Logger
	// OSSSetting is a pointer to a setting.OSSSettingS
	OSSSetting            *setting.OSSSettingS
)