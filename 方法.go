package main

import "fmt"

/*
go语言不是面向对象的语言
但是go里面的结构体嵌套和匿名字段可以实现继承
而且go语言的方法和函数是不一样的概念
操作结构体的话要用指针

函数：所有人都可以调用，不可以重名
func  方法名（方法参数）（返回类型）{
}
//只有接收者才可以调用，可以重名
func (接收者类型结构体) 方法名（方法参数）（返回类型）{
}
*/

type Cat struct {
	name string
	age  int
}

func (cat Cat) eat() {
	fmt.Println("cat eat")
}

type Dog struct {
	name string
	age  int
}

func (dog Dog) eat() {
	fmt.Println("dog eat")
}

func main() {
	cat := Cat{"c1", 5}
	cat.eat() //cat eat
	dog := Dog{"d1", 7}
	dog.eat() //dog eat
}
