package serve

import (
	"encoding/json"
	"fmt"
	"log"
	"msg_queue/db"
	"msg_queue/queue"
	"sync"
)

//以后好处理
var R []*queue.RabbitMQ

var PayTickets sync.Map //一个电影id对应一个

//这里是票的数量  默认读取多少
//默认是票的总数量
const Times = 10

type PayTicket struct {
	Tickets
	Left  int //剩余数量
	Msg   chan queue.Message
	PayCh chan int   //购买数量
	Lock  sync.Mutex //互斥锁
	Out   bool
}

//构造函数
func NewPayTicket(t *db.Tickets) *PayTicket {
	ticket := Tickets{t.MovieId, t.MovieName, t.WatchTime, t.WatchPlace, t.Num}

	p := &PayTicket{
		Tickets: ticket,
		Left:    ticket.Num,
		Msg:     make(chan queue.Message, 10),
		PayCh:   make(chan int),
	}
	p.Out = false
	return p
}

//初始化票
func InitMap() {
	_tickets, err := db.GetAllTickets()

	if err != nil {
		log.Println("get all tickets error")
		return
	}
	for _, v := range _tickets {

		PayTickets.Store(v.MovieId, NewPayTicket(&v))
	}

}

func (p *PayTicket) Consume() *queue.RabbitMQ {
	return queue.NewRabbitMQ("amqp://guest:guest@localhost:5672/", "", p.MovieId)
}

//从queue从获取订单
func (p *PayTicket) GetOrder() {
	//进行读取
	q := p.Consume()
	R = append(R, q)
	go q.ConsumeQueue()
	go func() {
		//这里是循环100ci
		for i := 0; i < p.Num; i++ {
			select {
			case m := <-q.AllMsg:
				p.Msg <- m


			}
		}
		p.Out = true
	}()

}

//从syncMap里取值
func getTicket(moveId string) *PayTicket {
	load, ok := PayTickets.Load(moveId)
	if !ok {
		log.Println("get ticket error")
		return nil
	}
	t := load.(*PayTicket)
	return t
}

//获取queue中的信息
func (p *PayTicket) GetMsg() {
	for {
		select {
		case order := <-p.Msg:
			var o Order
			fmt.Println("get order", string(order.Msg))
			if err := json.Unmarshal(order.Msg, &o); err != nil {
				log.Println("Unmarshal error", err)
				return
			}
			o.HappenTime = order.NowTime.Format("2006-01-02 15:04:05")
			p := getTicket(o.MovieId)

			o.MakeOrders()
			p.sellTickets()


		}

	}
}

func (p *PayTicket) sellTickets() {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	p.decTickets()
}

//减少票
func (p *PayTicket) decTickets() {
	p.PayCh <- 1 //抢票是一张一张的
}

//监听并扣除
func (p *PayTicket) Monitor() {
	for {
		select {
		case <-p.PayCh:
			p.Left -= 1

		}

	}
}

//开始服务
func StartServer() {
	InitMap()
	PayTickets.Range(func(k, v interface{}) bool {
		p := v.(*PayTicket)
		go p.Monitor()
		go p.GetOrder()
		go p.GetMsg()
		return true
	})

}
