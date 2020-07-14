package queue

import (
	"github.com/streadway/amqp"
	"log"
	"sync"
	"time"
)

type Message struct {
	Msg     []byte
	NowTime time.Time
}

//一个rabbieM
type RabbitMQ struct {
	Con       *amqp.Connection
	Ch        *amqp.Channel
	Url       string
	QueueName string
	Exchange  string
	Key       string
	AllMsg    chan Message
	Done      chan int
	Lock      sync.Mutex
}

//构造函数
func NewRabbitMQ(url string, exchange string, name string) (*RabbitMQ) {
	rabbitMQ := &RabbitMQ{
		Url:       url,
		QueueName: name,
		Exchange:  exchange,
		AllMsg:    make(chan Message),
		Done:      make(chan int),
	}
	var err error
	rabbitMQ.Con, err = amqp.Dial(rabbitMQ.Url)
	if err != nil {
		log.Println("connect error")
		return nil
	}
	rabbitMQ.Ch, err = rabbitMQ.Con.Channel()
	if err != nil {
		log.Println("chan error")
		return nil
	}
	return rabbitMQ
}

//关闭chan
func (r *RabbitMQ) destory() {
	r.Ch.Close()
	r.Con.Close()
}
