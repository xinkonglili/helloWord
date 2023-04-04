package main

import "fmt"

//panic: assignment to entry in nil map:因为没有初始化，map不像array和基础类型在你定义就会给你初始化一个默认值
//初始化map
func main() {
	var map0 map[int]string                            //只初始化，没赋值,未创建，为nil
	map1 := make(map[int]string)                       //只初始化，没赋值，,创建了，不为nil
	var map2 = map[string]int{"Go": 100, "Python": 99} //map[Go:100 Python:99]

	//map0[0] = "222"  无法赋值,因为没有给map0分配内存
	map1[1] = "看了没有" //make创建可以赋值
	map1[1] = "没有"     //map[1:没有]，相同的键，会采取覆盖的方式

	fmt.Println(map0)
	fmt.Println(map1) //map[1:看了没有]
	fmt.Println(map2)

	fmt.Println(map0 == nil) //true
	fmt.Println(map1 == nil) //false

}
