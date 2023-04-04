package main

import "fmt"

func main() {
	//常规的数组初始化
	var num01 [2]int
	fmt.Println(len(num01))
	fmt.Println(cap(num01))

	//使用 :=同时初始化定义
	num02 := [2]string{"nihao", "woyehao"}
	fmt.Println(num02)

	//定义变长的数组
	num03 := [...]string{"yyyy", "uuiui", "jfjdfkl", "oopop"}
	fmt.Println(num03)
	fmt.Println(len(num03))
	fmt.Println(cap(num03))

	//给指定的下标赋值,注意不要数组越界
	num04 := [4]int{0: 20, 3: 39} //[20 0 0 39]
	fmt.Println(num04)

}
