package rabbitmq

import (
	"context"
	"github.com/Penglq/QLog"
	"github.com/streadway/amqp"
	"git/zhuandd/account-service/utils"
	"sync"
	"time"
)

// func (p *MessageQueueConnection) GetChannel() {
// 	var err error
// 	p.Ch, err = p.Conn.Channel()
// 	if err != nil {
// 		//log.GetLogger().Alert("NewMessageQueueChannel", err.Error())
// 		panic(err)
// 	}
// 	return
// }

func (p *MessageQueueConnection) Get(ctx context.Context, ch chan int, wg *sync.WaitGroup, queueName string,
	fn func(ctx context.Context, msg amqp.Delivery, ch chan int, wg *sync.WaitGroup, mqch *amqp.Channel) error) (err error) {
	var (
		msg  amqp.Delivery
		ok   bool
		mqch *amqp.Channel
	)

	mqch, err = p.Conn.Channel()

	if err != nil {
		time.Sleep(p.SleepTime)
		return
	}

	msg, ok, err = mqch.Get(queueName, false)
	if err != nil {
		mqch.Close()
		QLog.GetLogger().Error("traceId", utils.GetTraceIdFromCTX(ctx), "方法名称", "Get", "取数据错误", err)
		time.Sleep(p.SleepTime)
		return
	} else if ok {
		wg.Add(1)
		ch <- 1
		go fn(ctx, msg, ch, wg, mqch)
	} else {
		mqch.Close()
		time.Sleep(p.SleepTime)
		return
	}
	return
}

func (p *MessageQueueConnection) GetMore(ctx context.Context, queueName string,
	callback func(ctx context.Context, msg amqp.Delivery) error) (err error) {
	var (
		msg  amqp.Delivery
		ok   bool
		mqch *amqp.Channel
	)

	mqch, err = p.Conn.Channel()
	if err != nil {
		time.Sleep(p.SleepTime)
		return
	}

	var d []amqp.Delivery
	for i := 0; i < 30; i++ {
		msg, ok, err = mqch.Get(queueName, false)
		if err != nil {
			// QLog.GetLogger().Alert("MessageQueue", fmt.Sprintf("Get数据报错，%s", err.Error()), "queue", queueName)
			break
		}
		if ok {
			if tmp, ok := msg.Headers[CreatedAt]; !ok || tmp == 0 {
				msg.Headers = amqp.Table{CreatedAt: time.Now().Unix()}
			}
			d = append(d, msg)
		} else {
			break
		}
	}

	for _, msg := range d {
		p.handleMsg(ctx, queueName, msg, callback)
	}
	mqch.Close()
	return
}

func (p *MessageQueueConnection) Write(exchange, routingKey string, body []byte, headers amqp.Table) (err error) {
	var mqch *amqp.Channel
	mqch, err = p.Conn.Channel()
	defer mqch.Close()

	err = mqch.Publish(
		exchange,   // exchange 可以传空
		routingKey, // routing key 可以队列名
		true,       // mandatory
		false,      // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Headers:      headers,
			Body:         body,
		})
	return
}

// 将当前时间写入
func (p *MessageQueueConnection) WriteData(ctx context.Context, exchange, routingKey string, msg amqp.Delivery) error {
	if tmp, ok := msg.Headers[CreatedAt]; !ok || tmp == 0 {
		msg.Headers = amqp.Table{CreatedAt: time.Now().Unix()}
	}
	QLog.GetLogger().Info(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", "WriteData", "action", "写入信息", "body", string(msg.Body), "header", msg.Headers)
	return p.Write(exchange, routingKey, msg.Body, msg.Headers)
}

func (p *MessageQueueConnection) handleMsg(ctx context.Context, queueName string, msg amqp.Delivery,
	callback func(ctx context.Context, msg amqp.Delivery) error) {
	err := callback(ctx, msg)
	if err != nil {
		err = p.WriteData(ctx, "", queueName, msg)
		if err != nil {
			QLog.GetLogger().Info(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", "callback", "error", err, "body", string(msg.Body))
		}
		if err = confirmOne(ctx, msg); err != nil {
			return
		}
	} else {
		confirmOne(ctx, msg)
	}
}

func confirmOne(ctx context.Context, msg amqp.Delivery) (err error) {
	if err = msg.Ack(false); err != nil {
		QLog.GetLogger().Info(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", "callback", "action", "消息确认", "error:", err, "body", string(msg.Body))
		return
	}
	return
}
