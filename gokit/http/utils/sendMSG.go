package utils

import (
	"git/miniTools/data-service/config"
	"github.com/penglq/QLog"
	"git/pkg/yrdSendMail"
	"log"
)

func SendEmail(toUsers []string, toCC []string, title, content string) {
	// 邮件配置
	var conf yrdSendMail.MailConf
	conf.ApiUrl = config.GetGlobalConfig().Email.ApiUrl
	conf.OrgNo = config.GetGlobalConfig().Email.OrgNo
	conf.AuthCode = config.GetGlobalConfig().Email.AuthCode
	conf.TplId = config.GetGlobalConfig().Email.TplId
	mailService := yrdSendMail.Init(conf)

	mailService.SetSubject(title)
	mailService.SetContent(content)
	mailService.SetAddress(toUsers)
	mailService.SetCc(toCC)

	resp, err := mailService.Send()
	if err != nil {
		QLog.GetLogger().Info("SendEmail Err", err)
		log.Printf(err.Error())
		return
	}
	QLog.GetLogger().Info("SendEmail Resp", resp)
	log.Println(resp)
}
