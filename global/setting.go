package global

import (
	"github.com/DrScorpio/ri-blog/pkg/logger"
	"github.com/DrScorpio/ri-blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.ApplicationSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	Logger          *logger.Logger
)
