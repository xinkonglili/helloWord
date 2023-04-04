package main

import (
	"fmt"
)

func main() {
	// 创建一个空的 map，用来存储字符串和整数类型的键值对
	myMap := make(map[string]int)

	// 创建一个空的 channel，用来传递整数类型的值
	myChannel := make(chan int)

	// 创建一个空的 channel，用来同步协程
	done := make(chan bool)

	// 启动一个协程，用来向 channel 中发送数据
	go func() {
		for i := 0; i < 5; i++ {
			myChannel <- i // 将整数 i 发送到 channel 中
		}
		close(myChannel) // 关闭 channel
	}()

	// 启动一个协程，用来从 channel 中接收数据，并将其存储到 map 中
	go func() {
		for value := range myChannel {
			myMap[fmt.Sprintf("key%d", value)] = value // 将整数 value 存储到 map 中
		}
		done <- true // 通知主协程协程已经完成任务
	}()

	// 等待协程完成任务
	//当执行到<- done语句时，程序会阻塞当前的goroutine，直到从done通道中接收到一个值为止
	<-done

	// 打印 map 中的内容
	for key, value := range myMap {
		fmt.Printf("%s: %d\n", key, value) //每次运行的结果可能不同
	}
}
