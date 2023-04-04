package main

import "fmt"

/*
实现继承的2种方式--结构体嵌套：
1、使用匿名变量，来提升字段  is-a的关系 a（父类） b（子类） b可以直接访问父类的属性
2、聚合关系，是has-a的关系 a（父类） b（子类） b不可以直接访问父类的属性
*/
//一个目录下面不能有同名的结构体
//一个目录下面的package要同名

//定义一个子类的结构体
type Student1 struct {
	name string
	age  int
}

//定义一个父类的结构体
type Teacher1 struct {
	Student1
	subject string //如果要在其他文件访问，首字母就必须要大写，在同一个文件，小写也可以
}

func main() {
	s1 := Student1{"ok", 14}
	fmt.Println("ok name---", s1.name)

	t1 := Teacher1{Student1{name: "nijao", age: 13}, "niaho"} //方式1：构建Student1
	t1.name = "jinli"                                         //方式2：在这里可以直接点出来name和age，因为匿名变量被提升字段了
	t1.age = 23
	//t1.subject = "计算机"
	fmt.Println(t1)

}
