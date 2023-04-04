package main

import (
	"fmt"
	"os"
)

func main() {
	//打开文件
	openfile_a, err := os.Open("./a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(openfile_a.Name()) //./a.txt

	//defer 延迟执行关闭文件，等下面的代码执行过后再执行关闭文件的操作
	defer openfile_a.Close()

	//设置缓冲区读取内容,每次读取一个字符，切片会存储我们读到的东西
	bs := make([]byte, 1, 1024)
	//n:读取到的数量
	n, err := openfile_a.Read(bs)
	fmt.Println(n)          //1
	fmt.Println(string(bs)) //读出的内容是6   EOF读到文件的末尾
	fmt.Println(err)        //<nil>

}
