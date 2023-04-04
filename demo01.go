package main

import "fmt"

//每个程序只有一个main函数的启动入口
func main() {
	fmt.Println("hello go")
	//定义变量，不用加冒号
	var name string = "jinli"
	//修改变量
	name = "wjinli"
	fmt.Println(name)
}
