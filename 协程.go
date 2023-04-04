package main

import "fmt"

//主goroutine
/*
1、先开启一个defer，来处理主goroutine的非异常的结束
2、启动一个专门用于后台清理垃圾的goroutine，并且设置gc可用标识
3、执行main包中所引用包下的init函数
4、执行main函数，执行main函数的时候还会检查是否发生panic，必要的时候处理异常
5、程序执行结束之后，主goroutine会终止自己的以及当下的进程的运行
*/

func main() {
	//并行是同时执行，多核cpu
	//交替执行，并发，交替的间隔非常之快，几乎感受不到
	//没开启携程，要先等hello这个方法执行完成，才能执行main函数里面的方法
	go hello()
	for i := 0; i < 55; i++ {
		fmt.Println("main----")
	}
}

func hello() {
	for i := 0; i < 20; i++ {
		fmt.Println("hello---")
	}
}
