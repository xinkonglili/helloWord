package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d is doing\n", id)
	defer wg.Done()
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	//开始10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	//等所有的协程完成之后
	wg.Wait()
	fmt.Println("所有的协程执行完成")

}
