package main

import "fmt"

//空接口是实现多态的一个很好的例子
type A interface {
}

func test11(a A) {
	fmt.Println("test11", a)
}

func test22(a interface{}) {
	fmt.Println("test22", a)
}

type Hand struct {
	name string
	size int
}

type Head struct {
	name string
	size int
}

func main() {
	//空接口：你给接口传什么类型的，他就代表什么类型的
	h1 := Head{name: "头", size: 12}
	h2 := Hand{name: "手", size: 14}
	h3 := "nijin"
	h4 := 18

	test11(h1) //test11 {头 12}
	test11(h3) //test11 nijin

	test22(h1) //test22 {头 12}
	test22(h3) //test22 nijin

	fmt.Println(h1) //{头 12}
	fmt.Println(h2) //{手 14}
	fmt.Println(h3) //nijin
	fmt.Println(h4) //18

	//空接口也可以传入map
	map1 := make(map[string]interface{})
	map1["jinli"] = 100
	map1["999"] = "nihao"
	map1["ok"] = Hand{ //可以存入结构体
		size: 18,
		name: "head",
	}
	fmt.Println("map--", map1)         //map[999:nihao jinli:100 ok:{head 18}]
	fmt.Println("map----", map1["ok"]) //map---- {head 18}

	//空接口也可以传入切片
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, h1, h2, h3, "23483", "jijijiji") //也可以自己加数据
	//slice-------- [{头 12} {手 14} nijin 23483 jijijiji]
	fmt.Println("slice--------", slice1)

}
