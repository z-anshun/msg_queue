package rsp

import "github.com/gin-gonic/gin"

//"code"的1开头为正确，0开头为错误

func OK(c *gin.Context,msg string){
	c.JSON(200,gin.H{
		"Code":"100",
		"Msg":msg,
	})
}

func JsonError(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"Code": "001",
		"Msg":  msg,
	})
}
//购买失败
func PayTicketError(c *gin.Context,msg string){
	c.JSON(200, gin.H{
		"Code": "002",
		"Msg":  msg,
	})
}
