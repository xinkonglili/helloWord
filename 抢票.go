package main

//出现临界资源抢夺的问题
import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

var ticket int = 10
var wg1 sync.WaitGroup

func main() {
	wg1.Add(4)

	go SaleTicket("窗口1")
	go SaleTicket("窗口2")
	go SaleTicket("窗口3")
	go SaleTicket("窗口4")

	fmt.Println("goroutine is ending")
	wg1.Wait()
	fmt.Println("限制解除")
	/*time.Sleep(time.Second * 5)*/

}

func SaleTicket(name string) {
	defer wg1.Done()
	for {
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Millisecond * 5)
			fmt.Println(name, "当前剩余票数：", ticket)
			ticket--
		} else {
			//操作完毕之后再去锁
			mutex.Unlock()
			fmt.Println("票已经售完")
			break
		}
		//操作完毕之后再去锁
		mutex.Unlock()
	}

}
