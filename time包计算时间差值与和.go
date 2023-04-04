package main

import (
	"fmt"
	"time"
)

func main() {
	//Add
	now := time.Now()
	laterTime := now.Add(time.Hour) //在现在的时间增加一个小时
	fmt.Println(now)
	fmt.Println(laterTime)

	//Sub
	sub := laterTime.Sub(now)
	fmt.Println("2个时间差值为：", sub) //2个时间差值为： 1h0m0s
	//Equal
	fmt.Println(now.Equal(laterTime)) //false

	//After 判断某个时间在指定的那个时间之后
	fmt.Println(now.After(laterTime)) //false
	//Before 判断某个时间在指定的那个时间之前
	fmt.Println(now.Before(laterTime)) //true

}
