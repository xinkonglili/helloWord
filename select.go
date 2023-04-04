package main

import (
	"fmt"
	"time"
)

/*
select 是 Go 中的一个控制结构，类似于 switch 语句。
select 语句只能用于通道操作，每个 case 必须是一个通道操作，要么是发送要么是接收。
select 语句会监听所有指定的通道上的操作，一旦其中一个通道准备好就会执行相应的代码块。
如果多个通道都准备好，那么 select 语句会随机选择一个通道执行。如果所有通道都没有准备好，那么执行 default 块中的代码。
*/
func main() {
	ch01 := make(chan int, 3)
	ch02 := make(chan int, 5)

	go func() {
		time.Sleep(time.Second)
		ch01 <- 100
	}()

	go func() {
		time.Sleep(time.Second)
		ch02 <- 200
	}()

	//每次运行的结果都不相同
	select {
	case data01 := <-ch01:
		fmt.Println("ch01", data01)
	case data02 := <-ch02:
		fmt.Println("ch02", data02)

	}

}
