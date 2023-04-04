package main

import "fmt"

func main() {
	var map3 map[int]string
	map3 = make(map[int]string, 99) //指定map的容量
	map3[4] = "100"
	map3[1] = "102"
	map3[0] = "999" //map[0] 值为  999

	fmt.Println(len(map3)) //3
	//当一个map变量被创建后，你可以指定map的容量，但是不可以在map上使用cap()方法
	//要获取map的容量，可以用len函数
	//fmt.Println(cap(map3))

	for ma := 0; ma < 5; ma++ {
		//可以用ok_idiom来判断key value是否存在
		value, ok := map3[ma]
		if ok {
			fmt.Printf("map[%d] 值为  %s", ma, value)
			fmt.Println()
		} else {
			fmt.Println("值不存在")
		}
	}

	delete(map3, 4)   //删除某个map3中的，下标为4的值
	fmt.Println(map3) //map[0:999 1:102]

}
