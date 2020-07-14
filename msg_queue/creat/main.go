package main

import (
	"github.com/streadway/amqp"
	"time"
)

type RabbitMQ struct {
	con       *amqp.Connection
	ch        *amqp.Channel
	url       string
	queueName string
	exchange  string
	key       string
}

func NewRabbitMQ(url string, exchange string, name string) *RabbitMQ {
	rabbitMQ := &RabbitMQ{
		url:       url,
		queueName: name,
		exchange:  exchange,
	}
	rabbitMQ.con, _ = amqp.Dial(rabbitMQ.url)
	rabbitMQ.ch, _ = rabbitMQ.con.Channel()
	return rabbitMQ
}

func main() {
	//初始化
	r := NewRabbitMQ("amqp://guest:guest@localhost:5672", "", "work")
	//指定队列   普通队列  topic（匹配） fanout （广播模型）
	r.ch.QueueDeclare(
		r.queueName,
		//是否持久化
		false,
		//是否为自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
	)


	r.ch.Publish(
		r.exchange,
		r.queueName,
		//如果为true，根据exchange类型和routkey规则，如果无法找到符合条件的队列，那么会把发送的消息返回给发送者
		false,
		//如果为true，当exchange发送消息到队列侯发现队列上没有绑定消费者，则会把消息发还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Timestamp:   time.Now(),
			Body:        []byte("test"),
		})
	defer r.ch.Close()
}
