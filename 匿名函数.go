package main

import (
	"fmt"
	"math"
)

func main() {
	f11() //我是f11
	f2 := f11
	f2() //我是f11

	/*f3 := func() {
		fmt.Println("我是f3")
	}
	f3()  //我是f3*/

	//写成匿名函数--函数的名字不用写
	func() {
		fmt.Println("我是f3") //我是f3
	}() //我是f3

	fmt.Println(hypro(3, 4)) //只能在函数内部使用

}

func f11() {
	fmt.Println("我是f11")
	hypro(3, 4)
}

//不定义函数名
func testNiMing() {
	func(a int, b string) int {
		return a
	}(20, "leo") //名字一样的话直接在后面省略掉函数的名字，因为本来也没有起名字，直接不用写
	testNM()
}

//将一个匿名函数赋值给一个变量
func testNM() {
	testNM := func(a int, b string) int {
		return a
	}
	testNM(20, "leo")
	testNiMing()
}

func hypro(x, y float64) float64 { //这里定义的函数只能在函数内部调用，不能在其他地方调用
	return math.Sqrt(x*x + y*y)
}
