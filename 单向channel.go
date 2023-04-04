package main

import (
	"fmt"
	"time"
)

//单向通道只能接收数据，或者只能发送数据，但是在函数中单独的去定义单向通道是没有意义的，主要是作为函数的参数传递的
func main() {
	//在channel上面加个方向就好了
	/*ch7 := make(chan<- int) //只可以写,定义单向通道的时候不需要加上变量的名称
	ch7 <- 1*/
	//data1 := <- ch7 //报错，无效的操作
	//fmt.Println("读出来的数据为：", data1)

	//ch8 := make(<-chan int) //只可以读
	//data := <-ch8
	//fmt.Println("读出来的数据为：", data)

	ch9 := make(chan int)
	go writeOnly(ch9)
	go readOnly(ch9)
	time.Sleep(time.Second)

}

//作为函数的参数传递
func writeOnly(ch chan<- int) {
	ch <- 1
}

func readOnly(ch <-chan int) {
	data := <-ch
	fmt.Println(data)
}
