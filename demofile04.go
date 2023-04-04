package main

//向文件中写入数据
import (
	"fmt"
	"os"
)

//使用OpenFile方法打开文件
func main() {
	//os.O_APPEND:直接追加在文件的后面，不会请调原来文件的内容os.O_RDONLY|os.O_WRONLY|os.O_APPEND
	//os.O_RDONLY|os.O_WRONLY：会清掉原来文件的内容
	file, err := os.OpenFile("./a.txt", os.O_RDONLY|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	fmt.Println(file.Name())
	if err != nil { //如果不是nil，就说明出错了
		fmt.Println(err)
		return
	}
	defer file.Close()
	//业务代码，写入数据
	bs := []byte{65, 66, 67} //ABC%&'666666888888
	n, err := file.Write(bs)
	n, err = file.WriteString("jinli888") //ABC%&'666666888888ABCABCjinli888
	fmt.Println(n)
	fmt.Println(err)
}
