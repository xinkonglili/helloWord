package main

import (
	"fmt"
	"time"
)

func main() {
	//自动定时器：Tick心脏跳动，按秒计时
	//tick := time.Tick(time.Second)
	/*
		C是一个chan time.Time类型的缓冲通道，一旦触及到期时间，定时器就会向自己的C字段发送一个time.Time类型的元素值
	*/
	/*fmt.Printf("%T\n", tick) //chan (通道)time.Time
	for t := range tick {
		fmt.Println(t)
	}*/

	//手动定时器
	for i := 0; i < 10; i++ {
		fmt.Println("===============")
		fmt.Println(time.Now())
		time.Sleep(time.Second)
	}

}
