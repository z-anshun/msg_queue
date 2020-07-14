package queue

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

//加入队列
func (r *RabbitMQ) Push(msg []byte) error {

	//指定队列   普通队列  topic（匹配） fanout （广播模型）
	_, err := r.Ch.QueueDeclare(
		r.QueueName,
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
	if err != nil {
		log.Println("queue declare error")
		return err
	}

	err = r.Ch.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true，根据exchange类型和routkey规则，如果无法找到符合条件的队列，那么会把发送的消息返回给发送者
		false,
		//如果为true，当exchange发送消息到队列侯发现队列上没有绑定消费者，则会把消息发还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Timestamp:   time.Now(),
			Body:        msg,
		})
	if err != nil {
		log.Println("publish error")
		return err
	}

	return nil
}
