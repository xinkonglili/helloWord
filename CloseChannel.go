package main

import (
	"fmt"
	"time"
)

func main() {
	ch2 := make(chan int)
	go testCh(ch2)

	for v := range ch2 {
		time.Sleep(time.Second)
		fmt.Println("read out到channel的数据为：", v)
	}

}

func testCh(ch chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		ch <- i
		fmt.Println("读入到channel的数据为：", i)
	}

	close(ch)
}
