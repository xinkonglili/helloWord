package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	/*
			Add() : 设置等待组中要执行的子goroutine的数量
			Wait(): 让主Goroutine等待其他的goroutine的执行完成，然后主goroutine再执行，Wait方法会阻塞
		    Done():让WaitGroup等待的Goroutine的数量减1
	*/
	wg.Add(2)
	go testWait1()
	go testWait2()

	fmt.Println("two goroutine is ending")
	wg.Wait()
	fmt.Println("阻塞解除")

}

func testWait1() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 3)
		fmt.Println("test1--", i)
	}
	wg.Done()
}

func testWait2() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 3)
		fmt.Println("test2--", i)
	}
	wg.Done()
}
