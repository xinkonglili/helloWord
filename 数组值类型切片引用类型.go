package main

import "fmt"

func main() {
	//数组：值传递
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	arr1 := arr
	fmt.Println(arr1)
	arr1[0] = 10
	//[1 2 3 4]
	//[10 2 3 4],数组2的改变，没有影响数组1
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println("----切片-----")
	//切片，引用传递
	arr3 := []int{6, 7, 8, 9}
	arr4 := arr3
	arr4[0] = 23
	//[23 7 8 9]
	//[23 7 8 9]  切片4改变，切片3也改变，切片4是由切片3拷贝过去的
	fmt.Println(arr3)
	fmt.Println(arr4)

}
