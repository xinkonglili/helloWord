package main

import (
	"fmt"
)

type In interface {
}

//使用switch来进行接口断言: in.(type)---in(接口变量)
func test33(in interface{}) {
	switch in.(type) {
	case string:
		fmt.Println("string类型的接口变量")
	case int:
		fmt.Println("int类型的接口变量")
	case Nihao:
		fmt.Println("Nihao类型的接口变量")
	case In:
		fmt.Println("I类型的接口变量")
	case nil:
		fmt.Println("nil类型的接口变量")
	default:
		fmt.Println("未知类型")
	}
}

type Nihao struct {
	name string
}

func main() {
	nihao1 := Nihao{name: "today"}
	a := 18
	b := "qqq"
	c := true
	test33(nihao1) //Nihao类型的接口变量
	test33(a)
	test33(b)
	test33(c) //I类型的接口变量

}

func assertInterface(i interface{}) {
	//okiodm
	i2, ok := i.(int)
	if ok {
		fmt.Println("接口断言成功")
		fmt.Println(i2)
	} else {
		fmt.Println("接口断言失败")
		fmt.Println(i2)
	}

}
