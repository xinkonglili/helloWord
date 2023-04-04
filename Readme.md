### 注意事项
#### 1、代码包包名和文件目录名，不要求一致，但是同一个目录下的文件的第一行声明的所属包，必须一致，比如文件目录叫hello，但是代码包包名可以声明为main
   <p>golang基础之context包的使用
     1、WithCancel：创建一个带有Clear()关闭函数的ctx 
     2、WithDeadline：创建一个带有超时时间点的ctx
     3、WithTimeout：创建一个有超时时间的ctx
     4、WithValue：创建了一个携带了参数的ctx
     5、Background：创建一个没有Deadline，没有Value，也不能被Cancel的emptyCtx
     6、TODO：也会创建一个emptyCtx
     7、Context的基础接口：Deadline(),Done(),Err(),Value()</p>





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