package main

//可以大幅度的提高读写效率，因为bufio包内嵌了缓冲区
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//给文件授权可读可写
	file, _ := os.OpenFile("D:\\softenvironment\\GoWorks\\src\\helloWord\\a.txt", os.O_RDWR, os.ModePerm)
	defer file.Close()

	//bufio包装
	reader := bufio.NewReader(file)

	buf := make([]byte, 1024, 1024)
	n, _ := reader.Read(buf) //读到的文件内容的个数
	fmt.Println(n)
	fmt.Println(string(buf[0:n]))

	//读取键盘输入NewReader
	inputReader := bufio.NewReader(os.Stdin)
	str, _ := inputReader.ReadString('\n') //buf.String()
	fmt.Println("读取到的键盘输入为：", str)

	//键盘写出
	writerReader := bufio.NewWriter(file)
	writeString, _ := writerReader.WriteString("hello" +
		"nihao000")
	fmt.Println(writeString)
	//由于缓冲区没满，无法写出，需要强制刷新
	writerReader.Flush()

}
