package db

import (
	"github.com/jinzhu/gorm"
	"log"
)

//电影id
//电影名
//观影时间
//观影地点
//电影票数量
type Tickets struct {
	gorm.Model
	MovieId    string `json:"movie_id" binding:"required"`
	MovieName  string `json:"movie_name" binding:"required"`
	WatchTime  string `json:"watch_time" binding:"required"`
	WatchPlace string `json:"watch_place" binding:"required"`
	Num        int    `json:"num" binding:"required"`
}

func GetAllTickets() (tickets []Tickets, err error) {
	if err := DB.Table("tickets").Find(&tickets).Error; err != nil {
		log.Println("get tickets error", err)
		return nil, err
	}
	return tickets, err
}
func AddTickets(t *Tickets) {
	if err := DB.Create(t).Error; err != nil {
		log.Println("add ticket error", err)
	}
}
