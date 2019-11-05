package net

import (
	"git/miniTools/data-service/config"
)

func initConfig() {
	config.InitConfig()
	config.InitLogger(config.AppName)
}
