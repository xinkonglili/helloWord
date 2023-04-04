package main

import (
	"fmt"
)

/*
1、调用函数的时候，形参指向实参的地址。
2、指针传递中，函数体内可以通过形参指针改变实参地址空间的内容。
运行结果：实参变量a的地址 0xc0000180a8
		实参变量b的地址 0xc0000180c0
		实参变量a的值 3
		实参变量b的值 7
		形参指针a的地址 0xc0000180a8
		形参指针b的地址 0xc0000180c0
		实参变量a的值 7
		实参变量b的值 3

可见实参a的地址和形参指针a的地址一样
*/

//a *int是个指针变量（形参：函数定义的时候）
func swap(a *int, b *int) { //在这个函数体内，可以通过形参指针改变实参地址空间的内容
	fmt.Println("形参指针a的地址", a)
	fmt.Println("形参指针b的地址", b)
	tmp := *a
	*a = *b
	*b = tmp
}

func main() {
	a := 3
	b := 7
	fmt.Println("实参变量a的地址", &a)
	fmt.Println("实参变量b的地址", &b)
	fmt.Println("实参变量a的值", a)
	fmt.Println("实参变量b的值", b)
	//（实参，函数调用的时候--是一个变量的地址）
	swap(&a, &b) //在进行函数调用的时候，使用&来修饰实参，表示是将该变量的地址作为参数传入函数，传给形参
	fmt.Println("实参变量a的值", a)
	fmt.Println("实参变量b的值", b)
}
