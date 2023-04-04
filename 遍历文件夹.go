package main

/*
D:\softenvironment\GoWorks\src\helloWord\c\d
D:\softenvironment\GoWorks\src\helloWord\c\d\e
D:\softenvironment\GoWorks\src\helloWord\c\d\e\a.txt
*/
import (
	"fmt"
	"os"
)

func main() {
	dir := "D:\\softenvironment\\GoWorks\\src\\helloWord\\c"
	ListDir(dir)
}

func ListDir(filePath string) {
	//ReadDir:读取文件夹
	fileDirInfo, _ := os.ReadDir(filePath)
	for _, file := range fileDirInfo { // _, file 返回index和value
		fileName := filePath + "\\" + file.Name()
		fmt.Println(fileName) //%s%s是个占位符

		//判断是不是文件夹，递归读取文件夹
		if file.IsDir() {
			ListDir(fileName)
		}
	}
}
