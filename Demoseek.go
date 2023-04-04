package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.OpenFile("D:\\softenvironment\\GoWorks\\src\\helloWord\\a.txt", os.O_RDWR, os.ModePerm)
	defer file.Close()

	//业务ABCDEF666666888888ABCABCjinli888hahahahhah
	file.Seek(2, io.SeekStart) //SeekStart：从文件的开头
	//创建一个缓冲区，把文件里面的内容读到缓冲区里面，再从缓冲区打印出来
	buf := []byte{0}
	file.Read(buf)
	fmt.Println(string(buf))

	//在文件后面追加内容
	file.Seek(0, io.SeekEnd) //SeekEnd：从文件的末尾
	file.WriteString("hahahahhah")

}
