package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

//断点续传
func main() {
	//传输源文件
	source := "D:\\softenvironment\\GoWorks\\src\\helloWord\\a.txt"
	//传输目的地
	destFile := "D:\\softenvironment\\GoWorks\\src\\helloWord\\c\\d\\e\\a.txt"
	//临时记录文件：记录上次读取source的时候读到的位置，也就是读到文件的第几个了，从1开始的，6个字母，就是6
	tempFile := "D:\\softenvironment\\GoWorks\\src\\helloWord\\1.txt"

	//与文件建立连接获取
	file1, _ := os.Open(source)
	file2, _ := os.OpenFile(destFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	file3, err := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	defer file1.Close()
	defer file2.Close()

	//读取文件记录的数据
	ret, err := file3.Seek(0, io.SeekStart) //从file3文件开始的位置，偏移量为0，记录数据的下标
	fmt.Println(ret)

	buf := make([]byte, 1024, 1024)
	n, _ := file3.Read(buf)      //从seek光标以后获取数据量
	countStr := string(buf[0:n]) //countStr：是string类型的，读取临时文件里面的内容
	fmt.Println("countStr", countStr)
	//count：是未来光标所在的位置
	count, _ := strconv.ParseInt(countStr, 10, 64) //count：是int类型的，10进制，默认大小64

	//seek：设置从未来光标所在的位置读取，SeekStart：是光标开始进来的位置
	seek, err := file1.Seek(count, io.SeekStart)
	fmt.Println(seek)
	r, err := file2.Seek(count, io.SeekStart)
	fmt.Println(r)

	bufData := make([]byte, 1024, 1024)
	total := int(count) //文件崩溃的位置

	for {
		readnum, err := file1.Read(bufData)
		fmt.Println(readnum)
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			file3.Close()
			//os.Remove(tempFile) //
			break
		}
		writeNum, err := file2.Write(bufData[:readnum])
		total = total + writeNum

		//将传输的进度存储到临时文件中
		file3.Seek(0, io.SeekStart)
		file3.WriteString(strconv.Itoa(total))

	}

}
