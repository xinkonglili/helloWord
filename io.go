package main

import (
	"fmt"
	"os"
)

//file类是封装os包中的。封装了文件的信息FileInfo(获取文件的信息，是个接口)，reader，Write
//类型 （-:文件，d:目录） 3组用户的权限（自己的，全局的，别人的）

func main() {
	//获取某个文件的状态
	fileInfo, err := os.Stat("D:\\softenvironment\\GoWorks\\src\\helloWord\\demofile04.go")
	if err != nil {
		return
	}
	fmt.Println(fileInfo.IsDir()) //false
	fmt.Println(fileInfo.Name())  //demofile04.go
	fmt.Println(fileInfo.Size())  //726字节数
	//&{32 {2861520921 30997286} {1616331640 30997289} {1616331640 30997289} 0 726}
	fmt.Println(fileInfo.Sys())
	fmt.Println("Mode----", fileInfo.Mode())       //mode（状况，状态）- rw- rw- rw- 文件，可读写，三个权限都是
	fmt.Println("ModTime----", fileInfo.ModTime()) // 2022-11-18 16:40:13.5992184 +0800 CST(文件的最后修改时间)

}
