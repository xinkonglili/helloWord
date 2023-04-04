package main

import "fmt"

func main() {
	myMap := make(map[string]int)
	myChannel := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 6; i++ {
			myChannel <- i
		}
		close(myChannel) //记得关闭，关闭并不影响读，只影响写
	}()

	go func() {
		for value := range myChannel {
			myMap[fmt.Sprintf("key%d", value)] = value
		}
		done <- true //协程完成任务
	}()

	//等待上面的协程完成任务
	<-done //当执行到<- done语句时，程序会阻塞当前的goroutine，直到从done通道中接收到一个值为止
	for k, v := range myMap {
		fmt.Printf("%s  %d\n", k, v)
	}
}
