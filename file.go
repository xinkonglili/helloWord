package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		type Writer struct {
			err error
			buf []byte
			n   int
			wr  io.Writer
		}*/
	writer := bufio.NewWriter(f)
	reader := bufio.NewReader(f)
	n := 0
	defer f.Close()
	for {
		n++
		str, err := reader.ReadString('\n')
		writer.WriteString(strconv.Itoa(n) + str)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	//在要写入的时候(Flush)，把光标移到最开始
	f.Seek(0, 0)
	writer.Flush()

}
