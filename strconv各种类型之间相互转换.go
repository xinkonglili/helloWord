package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "true"
	//ParseBool
	parseBool, _ := strconv.ParseBool(str)
	fmt.Printf("%T\n", parseBool) //bool
	fmt.Println("字符串转bool：", parseBool)

	//FormatBool
	formatBool := strconv.FormatBool(parseBool)
	fmt.Printf("%T\n", formatBool) //string
	fmt.Println("bool转字符串", formatBool)

	//ParseInt
	str1 := "100"
	i, _ := strconv.ParseInt(str1, 10, 64)
	fmt.Printf("%T\n", i) //int64
	fmt.Println("字符串转整数", i)

	//FormatInt
	formatInt := strconv.FormatInt(i, 10)
	fmt.Printf("%T\n", formatInt) //string
	fmt.Println("整数转字符串", formatInt)

}
