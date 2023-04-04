package main

import (
	"fmt"
	"time"
)

func main() {
	//延迟
	go spinner(100 * time.Millisecond)
	const n = 45
	FinN := Fin(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, FinN)
}
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}

	}
}

//1 1 2 3 5
func Fin(x int) int { //返回int64就报错
	if x < 2 {
		return x
	}
	return Fin(x-1) + Fin(x-2)
}
