package main

import (
	"fmt"
)

/*
运行结果：
start
recover from something went wrong
done
*/

func recoverDemo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover from", r)
		}
	}()

	fmt.Println("start")
	panic("something went wrong") // 抛出异常
	//程序执行到panic之后，就会发生panic，去执行defer后面的匿名函数，
	//程序的控制流转移到了defer中，recover可以恢复控制流，程序会在该匿名函数中继续执行，不会回到这里了
	fmt.Println("end") // 不会执行到这里
}

func main() {
	recoverDemo()
	fmt.Println("done")
}
