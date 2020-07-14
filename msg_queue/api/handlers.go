package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"msg_queue/db"
	"msg_queue/rsp"
	"msg_queue/serve"
)

//出售票
func PayTickets(c *gin.Context) {
	var order serve.Order
	if err := c.BindJSON(&order); err != nil {
		fmt.Println(order, err)
		rsp.PayTicketError(c, "pay error")
		return
	}
	if err := order.IsOut(); err != nil {
		rsp.PayTicketError(c, err.Error())
		return
	}
	if err := order.SendMsg(); err != nil {
		rsp.JsonError(c, "json error")
	}


	rsp.OK(c, "pay success")
	return

}
func AddMovie(c *gin.Context) {
	var movie serve.Tickets
	if err := c.BindJSON(&movie); err != nil {
		rsp.PayTicketError(c, "add error")
		return
	}
	db.AddTickets(&db.Tickets{
		MovieId:    movie.MovieId,
		MovieName:  movie.MovieName,
		WatchTime:  movie.WatchTime,
		WatchPlace: movie.WatchPlace,
		Num:        movie.Num,
	})
	c.JSON(200, gin.H{
		"data": movie,
	})
}
