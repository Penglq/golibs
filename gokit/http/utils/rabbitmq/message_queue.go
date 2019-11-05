package rabbitmq

import (
	"github.com/streadway/amqp"
	"git/miniTools/data-service/config"
	"github.com/penglq/QLog"
	"sync"
	"time"
)

var (
	jishu int
	l     int
	mqc   []*MessageQueueConnection
	mux   = sync.RWMutex{}
)

const CreatedAt = "created_at"

type MessageQueueConnection struct {
	Addr string
	Conn *amqp.Connection
	// Ch        *amqp.Channel
	L         *sync.RWMutex
	SleepTime time.Duration
}

func GetMQC() *MessageQueueConnection {
	mux.Lock()
	defer mux.Unlock()
	jishu++
	if jishu > 9999 {
		jishu = 0
	}
	return mqc[jishu%l]
}

func GetChannel() (*amqp.Channel, error) {
	return GetMQC().Conn.Channel()
}

func MessageQueueInit() {
	setAddr()
	l = len(mqc)
	for i := 0; i < l; i++ {
		_, err := mqc[i].connect()
		if err != nil {
			panic(err)
		}
	}
}

func setAddr() {
	l = len(config.GetGlobalConfig().Rabbitmq.Addrs)
	mux.Lock()
	defer mux.Unlock()
	var m []*MessageQueueConnection
	for i := 0; i < l; i++ {
		m = append(m, &MessageQueueConnection{Addr: config.GetGlobalConfig().Rabbitmq.Addrs[i]})
	}
	mqc = m
}

func (p *MessageQueueConnection) connect() (*MessageQueueConnection, error) {
	var err error
	QLog.GetLogger().Info("MessageQueue", "dial", "url", p.Addr)

	p.Conn, err = amqp.Dial(p.Addr)
	if err != nil {
		QLog.GetLogger().Info("MessageQueue", "重新连接......")
		time.Sleep(time.Second * 60)
		setAddr()
		p.connect()
		QLog.GetLogger().Alert("MessageQueue--->", "连接失败")
	} else {
		QLog.GetLogger().Info("MessageQueue", "连接成功......")
		// 监听连接关闭
		go p.confirmClose(p.Conn.NotifyClose(make(chan *amqp.Error, 1)))
	}
	return p, err
}

func (p *MessageQueueConnection) confirmClose(closeNotify <-chan *amqp.Error) {
	select {
	case confirmed := <-closeNotify:
		message := "客户端连接关闭"
		if confirmed.Server {
			message = "服务器端连接关闭"
		}
		p.connect()
		QLog.GetLogger().Info("MessageQueue", message, "code", confirmed.Code, "reason", confirmed.Reason)
	}
}

func (p *MessageQueueConnection) SetSleepTime(t time.Duration) {
	p.SleepTime = t
}
