package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	codn := sync.NewCond(&sync.Mutex{})
	/*codn.L.Lock()
	codn.Wait()
	codn.Wait()
	codn.Wait()
	codn.Wait()
	codn.L.Unlock()
	codn.Signal()    //通知wait完成一个，可能有好多个wait
	codn.Broadcast() //广播*/
	go func() {
		codn.L.Lock()
		fmt.Println("lock1")
		codn.Wait()
		fmt.Println("unlock1")
		codn.L.Unlock()

	}()

	go func() {
		codn.L.Lock()
		fmt.Println("lock2")
		codn.Wait()
		fmt.Println("unlock2")
		codn.L.Unlock()
	}()

	time.Sleep(1 * time.Second)
	codn.Signal()
	time.Sleep(1 * time.Second)
	codn.Signal()
	time.Sleep(1 * time.Second)

}
