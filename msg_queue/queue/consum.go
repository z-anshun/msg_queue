package queue

import (
	"log"
)

//消费队列
func (r *RabbitMQ) ConsumeQueue() {

	msgs, err := r.Ch.Consume(
		r.QueueName, // 队列名
		//用来区分多个消费者
		"",
		//是否自动应答
		true,
		//是否具有排他性
		false,
		//如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		//消息队列是否阻塞
		false,
		nil,
	)
	if err != nil {
		log.Println("creat consume-queue error ", err)

		return
	}

	go func() {
		for {
			select {
			case v := <-msgs:

				m := Message{
					Msg:     v.Body,
					NowTime: v.Timestamp,
				}
				r.AllMsg <- m
			}
		}
	}()
	for {
		select {
		case <-r.Done:
			return
		}
	}
}
