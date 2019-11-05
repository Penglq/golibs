package config

import (
	"github.com/penglq/QLog"
	"os"
)

func InitLogger(AppName string) {
	logLevel := QLog.INFO
	logger := QLog.GetLogger()
	isFile := true
	consulKey := os.Getenv("POD_CONSUL_TOKEN")
	if consulKey == "" {
		isFile = false
		logLevel = QLog.DEBUG
	}
	var config QLog.LoggerConfig
	config.LogLevel = logLevel
	config.IsConsole = true
	config.IsFile = isFile
	config.FilePath = LogPath
	config.Filename = AppName
	config.FileSuffix = "log"
	config.AlertConf = QLog.AlertApiConfig{
		AppId: GetGlobalConfig().Alert.AppId,
		URL:   GetGlobalConfig().Alert.URL,
	}
	logger.SetConfig(config)
	logger.SetTextPrefix("AppName", AppName)
}
