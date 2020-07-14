package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"msg_queue/api"
	"msg_queue/db"
	"msg_queue/serve"
	"msg_queue/timer"
)

func main() {
	e := gin.Default()
	api.SetRouter(e)
	//开始计时
	go timer.StartTimer()
	//初始化数据库
	db.InitDB()
	//开启服务
	serve.StartServer()
	fmt.Println("start..")
	e.Run(":8080")
}
