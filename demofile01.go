package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Stat("./a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fileInfo.IsDir())   //false
	fmt.Println(fileInfo.Name())    //a.txt
	fmt.Println(fileInfo.Mode())    //-rw-rw-rw-
	fmt.Println(fileInfo.ModTime()) //2022-11-18 14:45:32.9357891 +0800 CST
	fmt.Println(fileInfo.Size())    //18
	fmt.Println(fileInfo.Sys())     //&{32 {1217883562 30997273} {1529174083 30997273} {1529174083 30997273} 0 18}
}
