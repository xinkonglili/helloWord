package main

import "fmt"

//实现接口可以多写方法，但是不能少写方法
//接口没有关键字，实现了接口里面的方法就算实现接口了，但是必须实现接口中的所有方法
//go语言中的接口是非侵入式的，就算删掉接口，实现接口中的方法依然可以继续执行，只是变成了普通方法，只给那个结构体实现
//多态其实就是接口类型的传参

type Usb interface {
	IoInput()
	IoOutput()
}

type Mouse struct {
	name string
	size string
}

type KeyBoard struct {
	name string
	size string
}

func (mouse Mouse) IoInput() {
	fmt.Println("鼠标输入:", mouse.name, mouse.size)
}

func (mouse Mouse) IoOutput() {
	fmt.Println("鼠标输出:", mouse.name, mouse.size)
}

func (k1 KeyBoard) IoInput() {
	fmt.Println("鼠标输入:", k1.name, k1.size)
}

func (k1 KeyBoard) IoOutput() {
	fmt.Println("鼠标输出:", k1.name, k1.size)
}

//也可以写自己的接口
func (k1 KeyBoard) newMethod() {
	fmt.Println("自己定义的方法：", k1.name, k1.size)
}

func test(usb Usb) { //这里传接口
	usb.IoInput()
	usb.IoOutput()
}

func test1(usb Usb) { //这里传接口
	usb.IoInput()
	usb.IoOutput()
}

func main() {
	mouse1 := Mouse{"罗技", "18"}
	test(mouse1) //这里可以传实现接口的结构体，只要实现了接口的结构体，那么传接口和传接口的结构体是一样的
	board := KeyBoard{"lj", "28"}
	board.newMethod()
	test1(board)

	//实现的结构体可以赋值给接口，但是接口变量只能访问自己接口中定义的方法，实现类中的属性字段接口变量是访问不到的
	var usb Usb
	usb = mouse1
	fmt.Println(usb) //{罗技 18}
	//接口变量直接调用就好了，打印个锤子
	usb.IoOutput()            //鼠标输出: 罗技 18
	usb.IoInput()             //鼠标输入: 罗技 18
	fmt.Println(usb.IoInput)  //	0x46c160 打印的是地址
	fmt.Println(usb.IoOutput) //0x46c1c0 打印的是地址
}
