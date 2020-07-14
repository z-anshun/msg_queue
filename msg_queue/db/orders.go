package db

import (
	"github.com/jinzhu/gorm"
	"log"
)

//订单id
//订单发生时间
//订单人id
//电影id
type Order struct {
	gorm.Model
	OrderId    string
	HappenTime string  //这里是format后的
	UserId     string
	MovieId    string

}

func (order *Order) AddOrder() {
	if err := DB.Create(order).Error; err != nil {
		log.Println("add order error")
	}
}
