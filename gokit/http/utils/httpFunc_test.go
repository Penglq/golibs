package utils

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"git/miniTools/data-service/config"
	"git/miniTools/data-service/pkg/model"
	"testing"
	"time"
)

func TestHttpSend_HttpGet(t *testing.T) {
	config.InitConfig()
	config.InitLogger(config.AppName)
	ctx := context.Background()
	res, err := NewHttpClient(&ctx, UrlOptions("http://localhost:9999/"), TryOptions(2, time.Second*3)).HttpGet()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestRegexp(t *testing.T) {
	data := model.ToolsInsurance{}
	b, _ := json.Marshal(&data)
	t.Log(string(b))
}

func TestRandTime(t *testing.T) {
	for i := 0; i < 7; i++ {
		u, err := uuid.NewUUID()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(u.String())
	}
}
