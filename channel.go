package main

import "fmt"

//不要通过共享内存的方式来通信，应该通过通信来共享内存
//声明通道 var 通道名 chan 数据类型（通道里面装什么样的数据）
//创建通道 如果通道为nil就是不存在，就需要先创建通道
//通道名 = make(channel 要装的数据类型)

/*
发送和接收：data := <- a  //read from chan a
          a <- data  //write to chan a  箭头指向哪里就向哪里写数据
*/

/*
一个通道发送和接收数据默认是阻塞的，当一个数据被发送到通道中的时，在发送语句中被阻塞，直到另外一个Goroutine从该通道中读取
相对的当从通道读取数据的时候，读取被阻塞，直到一个goroutine将数据写入通道


channel本身就是同步的，意味着同一时间，只能有一条goroutine来操作，天然实现并发

如果goroutine在一个通道上面发送了数据，那么一定要有另外一个goroutine在来接收，否则将会出现死锁
如果goroutine在一个通道上面接收了数据，那么一定要有另外一个goroutine在来发送，否则将会出现死锁
如果创建了channel，但是没有goroutine使用，就会出现死锁
如果缓冲区满了，还往里面放的话也会deadlock

默认的channel只能放一个数据，设置容量之后就可以放入多个数据
*/

//先写再读
func main() {
	var ch chan bool
	ch = make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("我是：", i)
		}

		ch <- true //往通道写数据是没有等号的
	}()

	//读取通道中的数据
	data := <-ch
	fmt.Println("读取成功：", data) //读取成功： true   通道中就是一个true

}
