package main

import "fmt"

func main() {
	var ch1 chan int
	ch1 = make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("你好：", i)
			ch1 <- i //this is data: 0  管道是一次只能传一个值吗

		}
	}()

	data := <-ch1
	fmt.Println("this is data:", data)

}
