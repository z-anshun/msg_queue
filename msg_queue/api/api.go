package api

import "github.com/gin-gonic/gin"

func SetRouter(e *gin.Engine){
	e.POST("/payTickets",PayTickets)
	e.POST("/addMovie",AddMovie)

}
