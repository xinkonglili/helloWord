package main

import (
	"context"
	"fmt"
	"time"
)

/*
接收到值： 0
接收到值： 1
接收到值： 2
接收到值： 3
接收到值： 4
接收到值： 5
接收到值： 6
接收到值： 7
我结束了 <nil>

*/

//外层控制内层传参
func main() {
	ctx, clear := context.WithTimeout(context.Background(), time.Second*3) //确定需要上下文
	msg := make(chan int)
	go son(ctx, msg)
	//循环给message赋值
	for i := 0; i < 10; i++ {
		msg <- i
	}
	defer clear()
	time.Sleep(time.Second)
	fmt.Println("主进程结束")

}

//子函数
func son(ctx context.Context, msg chan int) {
	//定时器去接收
	t := time.Tick(time.Second) //1s调一下
	for _ = range t {
		select {
		case m := <-msg:
			fmt.Printf("接收到值： %d\n", m)
			// Done返回一个通道，该通道在工作完成时关闭
		case <-ctx.Done():
			fmt.Println("我结束了", ctx.Value("name"))
			return
		}
	}
}
