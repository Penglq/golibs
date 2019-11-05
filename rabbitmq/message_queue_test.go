package rabbitmq

import (
	"context"
	"github.com/streadway/amqp"
	"git/zhuandd/account-service/utils/config"
	"testing"
)

var addr string

func initConfig() {
	config.InitConfig()
	config.InitLogger("account-service")
	addr = config.GetGlobalConfig().Services.AccountService.Addr
	MessageQueueInit()
	// addr = "10.141.5.192:36406" //  开普勒
}
func TestMessageQueueConnection_Receive(t *testing.T) {
	initConfig()
	t.Log(config.GetGlobalConfig().Rabbitmq.Queue.WithdrawQuery.RoutingKey)
	err := GetMQC().WriteData(context.Background(), "", config.GetGlobalConfig().Rabbitmq.Queue.WithdrawQuery.RoutingKey, amqp.Delivery{Body: []byte("aaa")})
	if err != nil {
		t.Fatal("error", err)
	}
}
