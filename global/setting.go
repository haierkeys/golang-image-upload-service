package global

import (
	"github.com/haierspi/golang-image-upload-service/pkg/logger"
	"github.com/haierspi/golang-image-upload-service/pkg/setting"
)

var (
	// ServerSetting is a pointer to a setting.ServerSettingS
	ServerSetting *setting.ServerSettingS
	// AppSetting is a pointer to a setting.AppSettingS
	AppSetting *setting.AppSettingS
	// SecuritySetting is a pointer to a setting.SecuritySettingS
	SecuritySetting *setting.SecuritySettingS
	// EmailSetting is a pointer to a setting.EmailSettingS
	EmailSetting *setting.EmailSettingS
	// OSSSetting is a pointer to a setting.OSSSettingS
	OSSSetting *setting.OSSSettingS

	// Logger is a pointer to a logger.Logger
	Logger *logger.Logger
)
