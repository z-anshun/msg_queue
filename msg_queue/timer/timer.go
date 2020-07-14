package timer

import (
	"msg_queue/serve"
	"time"
)

func StartTimer() {

	defer func() {
		for _,v:=range serve.R{
			v.Done<-1
		}
	}()
	ticker := time.NewTicker(5 * time.Minute)
	select {
	case <-ticker.C:
		return
	}

}
