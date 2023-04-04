package main

import (
	"fmt"
	"math/big"
)

//SetUint64 把其他类型的整型存入int
func main() {
	big1 := new(big.Int).SetUint64(uint64(1000))
	fmt.Println("big1 is: ", big1)

	big2 := big1.Uint64()
	fmt.Println("big2 is: ", big2)
}
