package rabbitmq

import (
	"context"
	"github.com/streadway/amqp"
	"git/miniTools/data-service/config"
	"testing"
)

var addr string

func initConfig() {
	config.InitConfig()
	config.InitLogger("data-service")
	addr = ""
	MessageQueueInit()
	// addr = "10.141.5.192:36406" //  开普勒
}
func TestMessageQueueConnection_Receive(t *testing.T) {
	initConfig()
	t.Log()
	err := GetMQC().WriteData(context.Background(), "", "", amqp.Delivery{Body: []byte("aaa")})
	if err != nil {
		t.Fatal("error", err)
	}
}
