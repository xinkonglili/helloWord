package main

import "fmt"

func main() {
	arrs := [3][4]int{
		{5, 6, 3, 4},
		{2, 5, 6, 7},
		{1, 2, 3, 4},
	}
	//[5 6 3 4]
	//[2 5 6 7]
	//[1 2 3 4]
	for _, arr := range arrs {
		fmt.Println(arr)
	}
	fmt.Println("---------")
	for i := 0; i < 3; i++ {
		fmt.Println("开始打印第", i, "行")
		fmt.Printf("ksday  %d\n", i) //ksday  2  空格是在%前面加的

		for j := 0; j < 4; j++ {
			fmt.Println(arrs[i][j])
		}
	}
}
