package main

import (
	"fmt"
	"time"
)

//外层控制内层传参
func main() {
	flag := make(chan bool)
	msg := make(chan int)
	go son(flag, msg)
	//循环给message赋值
	for i := 0; i < 10; i++ {
		msg <- i
	}
	flag <- true
	time.Sleep(time.Second)
	fmt.Println("主进程结束")

}

//子函数
func son(flag chan bool, msg chan int) {
	//定时器去接收
	t := time.Tick(time.Second) //1s调一下
	for _ = range t {
		select {
		case m := <-msg:
			fmt.Printf("接收到值： %d\n", m)
		case <-flag: //检测有没有数据可以读
			//case flag <- : 检测有没有数据可以写
			fmt.Println("我结束了")
			return
		}
	}
}
