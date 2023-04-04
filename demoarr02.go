package main

import "fmt"

func main() {
	var nums = [4]int{78, 2, 3, 4}
	fmt.Println(nums[0])
	fmt.Println(nums[1])
	fmt.Println(nums[2])
	fmt.Println(nums[3])
	//[4]int:数组的长度加上数组的类型
	fmt.Printf("%T\n", nums)

	fmt.Println("--------------------")
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	fmt.Println("--------------------")
	for i, num := range nums {
		fmt.Println(i, num)
	}
	fmt.Println("--------------------")
	for i := range nums {
		fmt.Println(i)
	}

	fmt.Println("---------只打印数组中的值，忽略下标-----------")
	for _, v := range nums {
		fmt.Println(v)
	}

}
