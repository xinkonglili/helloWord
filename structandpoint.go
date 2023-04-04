package main

import "fmt"

func main() {
	/*
		a的值 10
		a的地址 0xc0000160a8
		b的值 0xc0000160a8
		b的地址 0xc000006028
	*/
	var b *int //声明一个指针变量
	//2种声明变量的方法
	var c int = 30
	a := 10
	fmt.Println(c)
	fmt.Println("a的值", a)
	fmt.Println("a的地址", &a)
	b = &a
	fmt.Println("b的值", b)
	fmt.Println("b的地址", &b)
	/*
		a的值 20
		a的地址 0xc0000160a8
	*/
	*b = 20
	fmt.Println("a的值", a) //a的值 20
	fmt.Println("a的地址", &a)

}
