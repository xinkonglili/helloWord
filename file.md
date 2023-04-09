
#### 
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	for {
		b := make([]byte, 12) //一个汉字三个字节
		n, err := f.Read(b)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b), n)
	}

}

```
```go
哈哈哈你 12
好吗 6
EOF

```
#### text.txt内容如下：哈哈哈你好吗怎么样了明天去玩嘛
#### f.Seek(-12, io.SeekEnd)
```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	z := 0
	defer f.Close()
	for {
		b := make([]byte, 12) //一个汉字三个字节
		n, err := f.Read(b)
		z++
		if z < 5 {
			f.Seek(-12, io.SeekEnd)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(b), n)
		}

	}

}

```
```go
哈哈哈你 12
天去玩嘛 12
天去玩嘛 12
天去玩嘛 12

```

#### 写入
```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	f.Seek(0, io.SeekEnd) //偏移量为0，从末尾开始写入
	f.Write([]byte("1234567"))
	if err != nil {
		fmt.Println(err)
		return
	}

}

```
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	/*
		// Reader实现了io的缓冲。读者对象.
		type Reader struct {
			buf          []byte
			rd           io.Reader // reader provided by the client
			r, w         int       // buf read and write positions
			err          error
			lastByte     int // last byte read for UnreadByte; -1 means invalid
			lastRuneSize int // size of last rune read for UnreadRune; -1 means invalid
		}*/
	/*
		func NewReader(rd io.Reader) *Reader {
			return NewReaderSize(rd, defaultBufSize)
		}
	*/
	reader := bufio.NewReader(f) //将文件变化为reader
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(str)
	}

}

```
```go
你好吗

怎么样了

明天去玩嘛

EOF

```

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	reader := bufio.NewReader(f) //将文件变化为reader
	for {
		str, isPrefix, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err.Error(), isPrefix) //前缀
			return
		}
		fmt.Println(string(str), isPrefix)
	}

}


```
```go
你好吗 false
怎么样了 false
明天去玩嘛 false
吃饭了吗 false
EOF false

```

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	//ReadDir返回[]DirEntry
	/*type DirEntry interface {
		Name() string
		IsDir() bool
		Type() FileMode
		Info() (FileInfo, error)
	}*/
	dInfo, err := os.ReadDir("./")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range dInfo {
		fmt.Println(v.Name())
		fmt.Println(v.IsDir())
		fmt.Println(v.Type())
		fmt.Println(v.Info())
	}
}

```

#### 这样写文件是没有变化的
```go
package main

import (
	"bufio"
	"fmt"
	"os"
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
	defer f.Close()
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		writer.WriteString(str + "1")
	}

}

```

#### bufio.NewWriter(f)  &&  bufio.NewReader(f)

```go
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

```
test.txt里面的内容变化
```go
1你好吗
2怎么样了
3明天去玩嘛
4吃饭了吗
```