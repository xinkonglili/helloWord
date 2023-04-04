package main

import "fmt"

func main() {

	type User struct {
		name     string
		age      int
		password string
	}

	var user1 User
	fmt.Println("----------", user1) //{0}
	user1.age = 18
	user1.name = "jinli"
	user1.password = "123"

	fmt.Println("----------", user1) //{jinli 18 123}
	fmt.Println(user1.age)           //18
	fmt.Println(user1.password)
	fmt.Println(user1.name)

	user2 := User{}
	user2.age = 17
	user2.name = "shdfjkj"
	user2.password = "689900"
	fmt.Println(user2)          //{shdfjkj 17 689900}
	fmt.Println(user2.name)     //shdfjkj
	fmt.Println(user2.password) //689900

	fmt.Println("-------------------------------")
	//第三种定义方式
	user3 := User{
		name:     "today",
		password: "123",
		age:      17,
	}
	fmt.Println(user3) //{today 17 123}

	//直接传参，字符串的顺序要一一对应
	user4 := User{"iii", 18, "90999"}
	fmt.Println(user4) //{iii 18 90999}

}
