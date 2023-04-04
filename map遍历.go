package main

import "fmt"

func main() {
	//map是无序集合，遍历出来的结果内容不一定按照顺序输出
	var map4 map[string]int
	map4 = make(map[string]int)
	map4["Go"] = 100
	map4["Python"] = 99
	map4["Java"] = 88
	fmt.Println("=========第一种遍历方式===========")
	for index, value := range map4 {
		//特殊化输出
		//fmt.Printf("%s:%d", index, value)
		//fmt.Println()

		//Go 100
		//Python 99
		//Java 88
		fmt.Println(index, value)

	}
	fmt.Println("=========第二种遍历方式===========")
	for k := range map4 {
		fmt.Println(k, map4[k])
	}
	fmt.Println(map4)
}
