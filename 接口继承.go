package main

import "fmt"

//接口嵌套
//接口可以多继承
type B interface {
	test2()
}

type D interface {
	test1()
}

//实现C的时候，也必须实现A,B的接口
type C interface {
	B
	D
	test3()
}

type Dogg struct {
	name string
}

//可以调用test3，但是不是接口类型的
func (dogg Dogg) test3() {
	fmt.Println("接口C--test3")
}

func (dogg Dogg) test1() {
	fmt.Println("接口D--test1")
}

func (dogg Dogg) test2() {
	fmt.Println("接口B--test2")
}

func main() {
	dogg := Dogg{
		name: "jinli",
	}
	dogg.test3()
	dogg.test1()
}
