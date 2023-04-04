package main

import (
	"fmt"
	"time"
)

func main() {
	a := 3
	go func() {
		a = 2
		fmt.Println("goroutine", a)

	}()

	a = 5
	time.Sleep(3 * time.Second) //保证goroutine上面的可以执行到
	//a变量的值已经被goroutine改过了，所以main 2
	fmt.Println("main", a)
}
