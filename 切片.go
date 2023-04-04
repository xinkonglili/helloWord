package main

import "fmt"

//注意一个切片追加，是针对当前的切片来说的
func main() {
	//长度为0，容量为10的slice，slice中的元素是int
	s1 := make([]int, 0, 9)
	fmt.Println(s1) //[]
	fmt.Println("--------")
	//append:是对同一个切片使用的，追加到同一个切片的
	s1 = append(s1, 1, 2, 4) //[1 2 4]
	fmt.Println(s1)

	//可以在一个切片上面扩容其他切片
	s2 := make([]int, 3, 4)
	s1 = append(s1, s2...) // s2...:取切片里面的值
	fmt.Println(s1)        //[1 2 4 0 0 0]

	//2、切片的遍历
	for i := 0; i < len(s1); i++ {
		fmt.Println(i) //这是打印下标
		fmt.Println("-------")
		fmt.Println(s1[i]) //这是打印值
	}
	fmt.Println("-------")
	for _, i2 := range s1 {
		fmt.Println(i2)
	}

}
