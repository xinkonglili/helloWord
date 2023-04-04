package main

import "fmt"

type current int

const (
	//先是变量名，后面才是变量的类型
	USD current = iota
	RMB
	EUA
)

func main() {
	symbol := [...]string{USD: "$", RMB: "￥", EUA: "^"}
	fmt.Println(RMB, symbol[RMB]) //1 ￥
	fmt.Println(USD, symbol[USD]) //0 $
	fmt.Println(EUA, symbol[EUA]) //2 ^

}
