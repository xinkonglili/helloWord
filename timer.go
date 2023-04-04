package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个定时器
	timer := time.NewTimer(time.Second * 3)
	//当前时间
	//timer.C,时间通道，这个通道中存放的值
	//就是我们定义的时候，
	fmt.Println(time.Now())
	timeChan := timer.C
	fmt.Println(<-timeChan)
	timer.Stop() //手动停止定时器
	//
}
