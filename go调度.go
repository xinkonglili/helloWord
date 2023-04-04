package main

import (
	"fmt"
	"runtime"
)

func main() {

	go func() {
		for i := 0; i < 15; i++ {
			fmt.Println("goroutine")
		}
	}()

	for i := 0; i < 15; i++ {
		//让出时间片，让别的goroutine先执行，但是不一定可以成功
		runtime.Gosched()
		fmt.Println("main---")
	}

}
