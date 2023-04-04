package main

import "fmt"

func main() {
	point := f1()
	fmt.Printf("point类型 %T\n", point) //*[4]int

	fmt.Println("point地址: ", &point) //0xc0000ce018

	fmt.Println("point指向的地址的全部数据: ", *point) //[1 2 3 4]

	fmt.Println("point指向的地址的第一个数据: ", point[0]) //1
	fmt.Println("----------------")
	fmt.Println(point) //&[1 2 3 4]

	a := 10
	f2(&a)
	fmt.Println("a的值为：", a) //a的值为： 100
}

//指针作为返回的参数类型
func f1() *[4]int { //存储的内容要一一对应
	array0 := [4]int{1, 2, 3, 4}
	return &array0
}

//指针作为参数
func f2(ptr *int) {
	*ptr = 100 //利用指针来修改变量
}
