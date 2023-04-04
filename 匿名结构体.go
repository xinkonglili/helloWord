package main

import (
	"fmt"
)

type Student struct {
	name     string
	password string
	age      int
}

//匿名定义结构体字段
type Teacher struct {
	string
	int
}

//结构体嵌套
type Address struct {
	addressId   int
	addressTime string
}
type Customer struct {
	name  string
	age   int
	adder Address
}

func main() {
	//正常情况
	student1 := Student{"jinli", "123", 16}
	fmt.Println(student1.name) //jinli

	//匿名定义结构体
	s2 := struct {
		name     string
		password string
	}{
		name:     "0000",
		password: "9999",
	}

	fmt.Println("匿名定义的结构体：", s2) // {0000 9999}
	fmt.Println("匿名定义的结构体：", s2.name)

	//结构体匿名字段
	t1 := Teacher{"6666", 27}
	fmt.Println("结构体匿名字段", t1.string) //6666
	fmt.Println("结构体匿名字段", t1.int)    //27

	//结构体嵌套
	c1 := Customer{}
	c1.name = "jinliC"
	c1.age = 23
	c1.adder = Address{ //这里要有之前结构体的名字
		addressId:   155,
		addressTime: "100",
	}

	fmt.Println("嵌套结构体：", c1)
	fmt.Println("嵌套结构体：", c1.adder.addressTime) //100
	fmt.Println("嵌套结构体：", c1.adder.addressId)   //155

}
