package utils

import (
	"bytes"
	"fmt"
	"git/miniTools/data-service/config"
	"git/miniTools/data-service/pkg/model"
	"html/template"
	"testing"
)

func init() {
	config.InitConfig()
	config.InitLogger("sendMail")
}

func TestSendEmail(t *testing.T) {
	var body bytes.Buffer
	tpl := template.New("email.html")
	tpl, _ = tpl.Parse(config.ConsultantTPL)
	err := tpl.Execute(&body, model.ToolsConsultant{
		ReservationTime: "2019",
		Name:            "小明",
		Mobile:          "1237126351",
		Sex:             1,
		Question1:       "adfasdfas",
		Answer1:         "2222222",
	})
	if err != nil {
		t.Log(err)
	}
	fmt.Println(body.String())
	// SendEmail([]string{"luqiangpeng@creditease.cn"}, []string{}, "hello", body.String())
}
