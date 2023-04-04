package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*	for i := range []int{1, 2, 3, 4} {
			rand.Seed(int64(i))
		}
	*/
	//用时间戳来随机
	rand.Seed(time.Now().Unix()) //如果要每次随机的数不一样，这里面的种子每次都要变的
	for i := 0; i < 10; i++ {
		num1 := rand.Intn(100) //生成[0,99]
		fmt.Println(num1)
	}

}
