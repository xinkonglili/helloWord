package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	ch4 := make(chan int)
	fmt.Println(len(ch4), cap(ch4)) //0 0

	ch5 := make(chan int, 5)
	fmt.Println(len(ch5), cap(ch5)) //0 5

	ch5 <- 1
	fmt.Println(len(ch5), cap(ch5)) //1 5

	//通道如果没满，可以一直往里面写数据，如果cap满了，数据还没有被使用，就会报错
	ch5 <- 2
	ch5 <- 3
	ch5 <- 4
	ch5 <- 5
	fmt.Println(len(ch5), cap(ch5)) //5 5
	//ch5 <- 6                        //fatal error: all goroutines are asleep - deadlock!
	//int 类型转 字符串用itoa

	ch6 := make(chan string, 4)
	go testCh3(ch6)
	for s := range ch6 { //read
		time.Sleep(time.Second)
		fmt.Println(s)
	}
}
func testCh3(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- "test-" + strconv.Itoa(i)
		fmt.Println("goroutine里面写入了数据" + strconv.Itoa(i))
	}

	close(ch)
}
