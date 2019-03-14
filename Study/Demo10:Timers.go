package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.NewTimer(time.Second*2)

	<- time1.C
	fmt.Println("t1 解放拉")

	//如果只是想等待多久再执行，也可以用time.Sleep。用timer的好处是，再触发之前，可以取消定时操作。
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}



}
