package main

import (
	"fmt"
	"io"
	"os"
)

//手动实现文件复制
func main() {
	source := "D:\\softenvironment\\GoWorks\\src\\helloWord\\a.txt"
	destination := "D:\\softenvironment\\GoWorks\\src\\helloWord\\c.txt"
	copy(source, destination, 1024)
}

func copy(source, destination string, bufSize int) {
	//输入文件   |文件内容复制不过去 要使用&
	sourceFile, err := os.OpenFile(source, os.O_RDONLY&os.O_WRONLY, os.ModePerm)
	//ourceFile, err := os.Open(source) //Open的权限高，可以对文件进行任何操作
	if err != nil {
		fmt.Println(err)
	}

	//延时关闭文件
	defer sourceFile.Close()

	//输出文件----先读到缓冲区
	destinationFile, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer destinationFile.Close()

	//先把文件输出到缓冲区,缓冲区大小1024，一次性读取1024个字节
	bs := make([]byte, bufSize)
	for {
		//读取
		n, err := sourceFile.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("文件复制完毕")
			break
		}
		//写出
		_, err1 := destinationFile.Write(bs[:n]) //每次缓冲区一满就读出来
		if err1 != nil {
			fmt.Println("写出失败", err1)
		}
	}
}
