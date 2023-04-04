package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
runtime包：获取系统信息，schedule调度让出时间片，让别的goroutine先执行
Goexit//终止当前的goroutine
*/

func main() {
	/*fmt.Println("获取goroot目录：", runtime.GOROOT()) //D:\Go\go
	fmt.Println("获取操作系统：", runtime.GOOS)         //windows
	fmt.Println("获取电脑几核cpu：", runtime.NumCPU())
	fmt.Println("获取执行的协程数：：", runtime.NumGoroutine())
	fmt.Println("获取go的版本：", runtime.Version())*/

	//这里是加了协程
	go func() {
		fmt.Println("start")
		test111()
		fmt.Println("end")
	}()

	time.Sleep(time.Second)
}

func test111() {
	//想成栈这种数据结构，先进后出
	defer fmt.Println("test defer")
	/*
		//运行结果
				start
				test defer
				end
			return
	*/

	/*
		//运行结果
			start
			test defer
	*/
	runtime.Goexit() //终止当前的goroutine
	fmt.Println("test")
}
