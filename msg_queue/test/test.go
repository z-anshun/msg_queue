package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"msg_queue/serve"
	"net/http"
	"time"
)

func main() {
	for i := 0; i < 10000; i++ {
		go do()
	}
	time.Sleep(time.Second * 10)

}
func do() {
	url := "http://localhost:8080/payTickets"

	o := serve.Order{
		OrderId:    "1",
		HappenTime: "",
		UserId:     "22",
		MovieId:    "123",
	}
	json_o, _ := json.Marshal(o)
	//提交请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json_o))
	if err != nil {
		log.Println("gg ", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(body))
}
