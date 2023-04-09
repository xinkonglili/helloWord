### 注意事项
#### 1、代码包包名和文件目录名，不要求一致，但是同一个目录下的文件的第一行声明的所属包，必须一致，比如文件目录叫hello，但是代码包包名可以声明为main
    tcp和http的区别:http是一种网络协议的规范，tcp是传输层的一种传输协议，tcp相当于一个网络库，会提供很多对外的接口，比如send，close，传输buf，这些buf可以是json格式，xml格式，http是无状态的协议，单次连接，连接之后就断开了，tcp要经历3次握手，4次挥手才会断开连接
#### golang基础之context包的使用
     1、WithCancel：创建一个带有Clear()关闭函数的ctx <br>
     2、WithDeadline：创建一个带有超时时间点的ctx  <br>
     3、WithTimeout：创建一个有超时时间的ctx
     4、WithValue：创建了一个携带了参数的ctx
     5、Background：创建一个没有Deadline，没有Value，也不能被Cancel的emptyCtx,确定需要使用上下文的时候使用
     6、TODO：也会创建一个emptyCtx，不确定使用上下文的时候使用
     7、Context的基础接口：Deadline(),Done(),Err(),Value()
     emptyCtx：实现了Context接口，它是个空的context，它永远不会被取消，没有值，没有deadline。其主要作为context.Background()和context.TODO()返回这种根context或者不做任何操作的context。
    如果用父子关系来理解，emptyCtx就是用来创建父context。
#### 应用场景
    Go 语言中的 context.Context 的主要作用还是在多个 Goroutine 组成的树中同步取消信号以减少对资源的消耗和占用
    在 Go 语言编写的服务器程序中，服务器通常要为每个 HTTP 请求创建一个 goroutine 以并发地处理业务。同时，这个 goroutine 也可能会创建更多的 goroutine 来访问数据库或者 RPC 服务。
    当这个请求超时或者被终止的时候，需要优雅地退出所有衍生的 goroutine，并释放资源。因此，我们需要一种机制来通知衍生 goroutine 请求已被取消。
#### 外层控制内层传参
    详解；为什么 time.Tick(time.Second) 可以放到range里面循环，博客地址:https://www.twle.cn/t/19246
    t := time.Newtimer(2time.Hour ) 初始化一个距离到期时间还有2个小时的一个日期，t是个指针类型的变量
    通过定时器的字段 C, 我们可以及时得知定时器到期的这个事件来临，C 是一个 chan time.Time 类型的缓冲通道，
    一旦触及到期时间，定时器就会向自己的 C 字段发送一个 time.Time 类型的元素值 
``` 
case <- t.C:
   do_some_thing()
```

```
    package main
    
    import (
        "fmt"
        "time"
    )
    
    //外层控制内层传参
    func main() {
        flag := make(chan bool)
        msg := make(chan int)
        go son(flag, msg)
        //循环给message赋值
        for i := 0; i < 10; i++ {
            msg <- i
        }
        flag <- true
        time.Sleep(time.Second)
        fmt.Println("主进程结束")
    
    }
    
    //子函数
    func son(flag chan bool, msg chan int) {
        //定时器去接收
        t := time.Tick(time.Second) //1s调一下
        for _ = range t {
            select {
            case m := <-msg:
                fmt.Printf("接收到值： %d\n", m)
            case <-flag: //检测有没有数据可以读
                //case flag <- : 检测有没有数据可以写
                fmt.Println("我结束了")
                return
            }
        }
    }
```

```go
package main

import (
	"context"
	"fmt"
	"time"
)

//外层控制内层传参
func main() {
	ctx, clear := context.WithCancel(context.Background()) //确定需要上下文    
	msg := make(chan int)
	go son(ctx, msg)
	//循环给message赋值
	for i := 0; i < 10; i++ {
		msg <- i
	}
	clear() //使上下文关闭
	time.Sleep(time.Second)
	fmt.Println("主进程结束")

}

//子函数
func son(ctx context.Context, msg chan int) {
	//定时器去接收
	t := time.Tick(time.Second) //1s调一下
	for _ = range t {
		select {
		case m := <-msg:
			fmt.Printf("接收到值： %d\n", m)
			// Done返回一个通道，该通道在工作完成时关闭
		case <-ctx.Done():
			fmt.Println("我结束了")
			return
		}
	}
}

```

```go
package main

import (
	"context"
	"fmt"
	"time"
)

//外层控制内层传参
func main() {
	ctx := context.WithValue(context.Background(), "name", "jinli")
	//这里ctx就携带了属性val和key这个属性
	ctx, clear := context.WithCancel(ctx) //确定需要上下文
	msg := make(chan int)
	go son(ctx, msg)
	//循环给message赋值
	for i := 0; i < 10; i++ {
		msg <- i
	}
	clear() //使上下文关闭
	time.Sleep(time.Second)
	fmt.Println("主进程结束")

}

//子函数
func son(ctx context.Context, msg chan int) {
	//定时器去接收
	t := time.Tick(time.Second) //1s调一下
	for _ = range t {
		select {
		case m := <-msg:
			fmt.Printf("接收到值： %d\n", m)
			// Done返回一个通道，该通道在工作完成时关闭
		case <-ctx.Done():
			fmt.Println("我结束了", ctx.Value("name"))
			return
		}
	}
}

```

```go
package main

import (
	"context"
	"fmt"
	"time"
)

/*
接收到值： 0
接收到值： 1
接收到值： 2
接收到值： 3
接收到值： 4
接收到值： 5
接收到值： 6
接收到值： 7
我结束了 <nil>

*/

//外层控制内层传参
func main() {
	//ctx, clear := context.WithTimeout(context.Background(), time.Second*3)  //3s之后协程结束
	ctx, clear := context.WithDeadline(context.Background(), time.Now().Add(time.Second*8)) //确定需要上下文
	msg := make(chan int)
	go son(ctx, msg)
	//循环给message赋值
	for i := 0; i < 10; i++ {
		msg <- i
	}
	defer clear()
	time.Sleep(time.Second)
	fmt.Println("主进程结束")

}

//子函数
func son(ctx context.Context, msg chan int) {
	//定时器去接收
	t := time.Tick(time.Second) //1s调一下
	for _ = range t {
		select {
		case m := <-msg:
			fmt.Printf("接收到值： %d\n", m)
			// Done返回一个通道，该通道在工作完成时关闭
		case <-ctx.Done():
			fmt.Println("我结束了", ctx.Value("name"))
			return
		}
	}
}

```