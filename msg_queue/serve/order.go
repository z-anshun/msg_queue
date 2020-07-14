package serve

import (
	"encoding/json"
	"errors"
	"msg_queue/db"
	"msg_queue/queue"
)

//订单id
//订单发生时间
//订单人id
//电影id
type Order struct {
	OrderId    string `json:"order_id" binding:"required"`
	HappenTime string `json:"happen_time"` //这里是format后的
	UserId     string `json:"user_id" binding:"required"`
	MovieId    string `json:"movie_id" binding:"required"`
}

func (o *Order) MakeOrders() {
	order := db.Order{
		OrderId:    o.OrderId,
		HappenTime: o.HappenTime,
		UserId:     o.UserId,
		MovieId:    o.MovieId,
	}
	order.AddOrder()
}
//判断是否还有
func (o *Order) IsOut() error {
	load, ok := PayTickets.Load(o.MovieId)
	if !ok {
		return errors.New("no movie")
	}
	p := load.(*PayTicket)

	if p.Out {
		return errors.New("the tickets is out")
	}
	return nil
}
//发送信息
func (o *Order)SendMsg()error{
	json_order, err := json.Marshal(o)
	if err != nil {
		return err
	}
	//以movieId做队列名
	r := queue.NewRabbitMQ("amqp://guest:guest@localhost:5672/", "", o.MovieId)
	r.Push(json_order)
	return nil
}