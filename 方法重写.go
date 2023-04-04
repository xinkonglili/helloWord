package main

//方法重写
import "fmt"

//父类
type Animal struct {
	name string
	age  int
}

//子类
type Dog1 struct {
	Animal
}

//子类
type Cat1 struct {
	Animal
	color string
}

//父类的eat方法
func (animal Animal) eat() {
	fmt.Println("animal---eat")
}

//子类重写父类方法
func (dog Dog1) eat() {
	fmt.Println("dog1----eat")
}

//子类可以有自己的方法
func (cat Cat1) miao() {
	fmt.Println("cat1----miao")
}

func main() {
	animal := Animal{name: "animal1", age: 29}
	animal.eat()
	dog1 := Dog1{Animal{"dog1", 18}}
	dog1.eat()
	cat1 := Cat1{Animal{name: "cat1", age: 17}, "red"}
	fmt.Println(cat1.color)
	//子类没有重写父类的方法，调用的就是父类的方法
	cat1.miao() //cat1----miao
	cat1.eat()  //animal---eat
}
