package main

import (
	"fmt"
)

func main() {
	//切片不用定义数组的长度，切片是引用类型
	var nums []int
	fmt.Printf("%T\n", nums) //[]int
	//一个切片在未初始化的时候默认为nil
	if nums == nil {
		fmt.Println("切片默认为nil") //切片默认为nil
	}
}
