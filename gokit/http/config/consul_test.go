package config

import (
	"testing"
	"time"
)

func TestInitConsul(t *testing.T) {
	consulToken = "e628e1b0-afc4-08b0-8e98-93b00b7ba0d8"
	consulKey = "operations.miniTools"
	consulUrl = "http://consul.yixinonline.org:8500"
	InitConsul()
	time.Sleep(time.Hour)
}

func TestKVGetConfig(t *testing.T) {
	appConfigs = &appConfig{}
	consulToken = "e628e1b0-afc4-08b0-8e98-93b00b7ba0d8"
	consulKey = "operations.miniTools"
	consulUrl = "http://consul.yixinonline.org:8500"
	KVGetConfig(consulQuestionKey, parseConsulQuestionConfig)
}
