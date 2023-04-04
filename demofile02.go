package main

/*
1、创建文件夹：Mkdir
2、删除文件夹：Remove
3、创建文件：Create
*/
import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("D:\\softenvironment\\GoWorks\\src\\helloWord\\b", os.ModePerm)
	if err != nil {
		fmt.Println(err) //Cannot create a file when that file already exists.
	}
	//创建文件夹出现的原因是上一层文件夹没有创建，所以找不到
	/*///*	err1 := os.Mkdir("D:\\softenvironment\\GoWorks\\src\\helloWord\\c\\d\\e", os.ModePerm)
	//	if err1 != nil {
	//		fmt.Println(err1) //The system cannot find the path specified.
	//	}
	//	fmt.Println("文件夹创建完毕")*/
	//创建文件夹出现的原因是上一层文件夹没有创建，所以找不到--solution:MkdirAll方法
	err1 := os.MkdirAll("D:\\softenvironment\\GoWorks\\src\\helloWord\\c\\d\\e", os.ModePerm)
	if err1 != nil {
		fmt.Println(err1) //The system cannot find the path specified.
	}
	fmt.Println("文件夹创建完毕")

	/*删除文件*/
	/*os.Remove("D:\\softenvironment\\GoWorks\\src\\helloWord\\b")
	err2 := os.Mkdir("D:\\softenvironment\\GoWorks\\src\\helloWord\\b", os.ModePerm)
	if err2 != nil {
		fmt.Println(err2) //Cannot create a file when that file already exists.
	}*/

	fileCreate, err1 := os.Create("D:\\softenvironment\\GoWorks\\src\\helloWord\\b.txt")
	fmt.Println(fileCreate.Name())
	fmt.Println(fileCreate.Chmod(os.ModePerm))
	if err1 != nil {
		fmt.Println(err1)
	}
}
