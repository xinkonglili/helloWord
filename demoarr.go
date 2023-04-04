package main

import "fmt"

/*
数组是具有相同类型的一组数据的有序集合
每个数据元素都可以通过下标来访问他们
*/
func main() {
	//var 变量名 [大小]变量类型，变量定义了就要使用
	var nums [4]int
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4

	//%T打印数据类型
	fmt.Printf("%T\n", nums) //[4]int
	//通过下标来获取数据元素
	fmt.Println(nums[0])
	fmt.Println(nums[1])
	fmt.Println(nums[2])
	fmt.Println(nums[3])

	//由于数组长度是固定的，所以数组的长度和容量是一样的
	fmt.Println(len(nums)) //数组长度 4
	fmt.Println(cap(nums)) //数组容量 4

	fmt.Println(nums) //[1 2 3 4]
	nums[0] = 100
	//可以直接打印数组
	fmt.Println(nums) //[100 2 3 4]

}
