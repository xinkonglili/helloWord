package main

import "fmt"

//数组是一个固定长度的，由特定类型的元素组成的序列
//数组可以由0个或者多个元素组成，由于数组的长度固定，所以很少的情况下会使用数组，一般使用slice
//slice是可以增长，收缩的动态序列

func main() {
	//定义数组
	var a [3]int     //默认值是0
	var b [50]string //默认值为空，不赋值，元素全为空
	//可以使用数组字面值的方法来初始化数组
	//var d [3]int = [3]int{1, 2, 3} //数组的长度最大为3

	//省略号--初始赋值多少，数组的长度就是多少
	d := [...]int{1, 2, 3} //省略号说明数值的长度随着你赋值多少而变化的,更改数组d写的方式
	for _, v := range d {
		fmt.Printf("%d \n", v)
	}
	fmt.Println("0000000000", a[0]) //0
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v) //0 0，1 0，2 0
	}
	//可以只是打印下标，也可以只是打印值
	for _, v := range a {
		fmt.Printf("%d \n", v)
	}

	for i, _ := range a {
		fmt.Printf("%d \n", i)
	}
	fmt.Println("1111111111--", b[0])
	fmt.Println(b[len(b)-1])
}
