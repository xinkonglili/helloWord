package main

import "fmt"

func main() {
	//拿到地址就拿到了数据，指针本身就是存的地址，所以拿到指针也就拿到了数据
	//数组指针
	arry0 := [4]int{1, 2, 3, 4} //普通数组
	var arry1 *[4]int
	var arry2 **[4]int
	arry1 = &arry0 //arry1是个指针变量
	arry2 = &arry1
	fmt.Println(arry0)                     //[1 2 3 4]
	fmt.Println("指针指向地址的内存里面的值", arry0[0]) //array0就是拿到array0这块区域上的
	fmt.Println("*操作", (*arry1)[0])        //1  拿到所有的数据，我只取一个
	fmt.Println("指针操作", arry1[0])          //1
	fmt.Println("===================================")
	fmt.Printf("arry0的地址： %p\n", &arry0) //变量本身的地址   0xc00000e180
	//arry1这种直接取出指针变量里面的值
	fmt.Printf("arry1里面存的数据为（也就是arry0的地址）： %p\n", arry1) //ayyy1里面存的值为arry0的地址  0xc00000e180
	fmt.Printf("array1的地址  %p\n", &arry1)                //arry1变量本身的地址 0xc000006028
	fmt.Println("arry0里面存储的数据", *arry1)                  // 这个地址上的所有数据都要 [1 2 3 4] 指针指向的变量地址里面的内容
	fmt.Println("==============----------============")
	fmt.Printf("arry2里面存的数据为（也就是arry1的地址）: %p\n", arry2) //0xc000006028
	fmt.Printf("array2的地址 %p\n", &arry2)                 //0xc000006030
	fmt.Printf("arry1里面存储的数据(arry0的地址) %p\n", *arry2)    //0xc00000e180
	fmt.Println("arry0里面存储的数据(arry0里面的数据)", **arry2)     // [1 2 3 4]

	//指针数组
	a := 10
	b := 20
	c := 30
	d := 40
	//第一种定义的方式
	var array3 [4]*int
	array3[0] = &a
	array3[1] = &b
	array3[2] = &c
	array3[3] = &d

	//第二种定义的方式
	array4 := [4]*int{&a, &b, &c, &d}
	fmt.Println(array3)    //[0xc0000160e0 0xc0000160e8 0xc0000160f0 0xc0000160f8]
	fmt.Println(array3[0]) //输出（a变量的地址）：&a
	*array3[0] = 70
	fmt.Println(array4)     //[0xc0000aa090 0xc0000aa098 0xc0000aa0a0 0xc0000aa0a8]
	fmt.Println(*array4[0]) //10
	*array4[0] = 500
	fmt.Println(*array4[0]) //500

}
