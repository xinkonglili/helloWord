package main

import "fmt"

func main() {
	var a, b int
	var pwd int = 2022
	fmt.Println("请输入密码：")
	fmt.Scan(&a)  //到a这里就停止了
	if a == pwd { //不用括号
		fmt.Println("请再次输入密码：")
		fmt.Scan(&b)
		if b == pwd {
			fmt.Println("登录成功!")
		}
	} else {
		fmt.Println("密码错误，登录失败")
	}
}
