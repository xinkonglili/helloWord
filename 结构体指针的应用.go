package main

//结构体是值传递，如果操作的话只能用指针操作
import (
	"fmt"
)

type User struct {
	name     string
	password string
	age      int
}

func main() {
	user1 := User{}
	user2 := user1
	user2.name = "jinli"
	fmt.Println("------", user1.name) //为空，说明是值传递

	var user3 User
	user3.age = 18
	user4 := user3
	user4.age = 23
	fmt.Println(user3.age) //18

	//定义结构体指针
	var user5 *User
	user5 = &user3
	user5.age = 5
	fmt.Println("使用结构体操作的user3：", user3.age) //结果为5

	//使用内置函数new出来的对象就是指针类型
	user6 := new(User)                   //指针类型
	fmt.Printf("new出来的内置对象：%T\n", user6) //*main.User
	user6 = &user3
	user6.age = 9
	fmt.Println("使用new函数创建出来的指针操作的user3：", user3.age) //9

}
