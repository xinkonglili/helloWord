### sync包的使用
    1、MUtex互斥锁
       - Lock()  写的时候 排斥其他的读锁和写锁
       - Unlock()

    2、RWMutex读写互斥锁
        - Lock()  写的时候 排斥其他的读锁和写锁
        - Unlock()
        - Rlock()  在读取的时候，不会排斥其他的读取锁，但是会排斥掉写锁
        - Runlock()

    3、Once
        - Once.Do(一个函数)这个方法无论被调用多少次，这里只会执行一次
    4、WaitGroup
        - Add(delta int)  设定需要执行Done多少次
        - Done() Done一次+1
        - Wait() 在达到Done的次数前一直阻塞
    5、Map一个并发字典
        - Store：写
        - Load：读
        - LoadOrStore:读的时候如果不存在，就塞进去
        - Range
        - Delete
#### 互斥锁
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	lock := &sync.Mutex{}
	go lockFun(lock)
	/*
		func lockFun(lock *sync.Mutex) {
			lock.Lock()
			fmt.Println("疯狂输出")
			time.Sleep(1 * time.Second)
			lock.Unlock()
		}*/
	go lockFun(lock)
	/*
		func lockFun(lock *sync.Mutex) {
			lock.Lock()
			fmt.Println("疯狂输出")
			time.Sleep(1 * time.Second)
			lock.Unlock()
		}*/
	go lockFun(lock)
	/*
		func lockFun(lock *sync.Mutex) {
			lock.Lock()
			fmt.Println("疯狂输出")
			time.Sleep(1 * time.Second)
			lock.Unlock()
		}*/
	go lockFun(lock)
	/*
		func lockFun(lock *sync.Mutex) {
			lock.Lock()
			fmt.Println("疯狂输出")
			time.Sleep(1 * time.Second)
			lock.Unlock()
		}*/
	time.Sleep(5 * time.Second) //主线程执行太快了，导致其他的goroutine还没开始执行，就退出了

}

func lockFun(lock *sync.Mutex) {
	lock.Lock()
	fmt.Println("疯狂输出")
	time.Sleep(1 * time.Second)
	lock.Unlock()
}

```

#### 读写锁
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	lock := &sync.RWMutex{}
	go lockFun(lock)
	go lockFun(lock)
	go readLockFun(lock)
	go readLockFun(lock)
	go readLockFun(lock)
	go readLockFun(lock)

	time.Sleep(5 * time.Second) //主线程执行太快了，导致其他的goroutine还没开始执行，就退出了

}

func lockFun(lock *sync.RWMutex) {
	lock.Lock() //写的时候 排斥其他的读锁和写锁
	fmt.Println("疯狂写入")
	time.Sleep(1 * time.Second)
	lock.Unlock()
}

/*
一起刷出来的
疯狂读入
疯狂读入
疯狂读入
疯狂读入*/
func readLockFun(lock *sync.RWMutex) {
	lock.RLock() //在读取的时候，不会排斥其他的读取锁，但是会排斥掉写锁
	fmt.Println("疯狂读入")
	time.Sleep(1 * time.Second)
	lock.RUnlock()
}

```

#### once
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	once := &sync.Once{}

	for i := 1; i < 10; i++ {
		once.Do(func() {
			fmt.Printf("执行%d", i)
		})

	}

}

```
```go
结果：
执行1
Process finished with the exit code 0****
```
#### 等待组，等待内部的协程通知你，他已经执行完成了
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		time.Sleep(5 * time.Second)
		wg.Done() //开始打血
		fmt.Println("掉血-1")
	}()

	go func() {
		time.Sleep(7 * time.Second)
		wg.Done() //开始打血
		fmt.Println("掉血-1")
	}()

	wg.Wait()
	fmt.Println("out")
}

```

```go
package main

import "fmt"

func main() {
	//这种二者不断的进行读写，导致程序崩溃
	m := make(map[int]int) //普通map
	go func() {
		for {
			m[1] = 1
		}

	}()

	go func() {
		for {
			fmt.Println(m[1])
		}

	}()
	for {

	}
}

```

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//这种map不会造成读写冲突
	m := &sync.Map{}
	go func() {
		for {
			m.Store(1, 1)
		}

	}()

	go func() {
		for {
			fmt.Println(m.Load(1))
		}

	}()
	time.Sleep(100)
}

```

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := &sync.Map{}
	m.Store(1, 1)
	m.Store(2, 2)
	m.Store(3, 3)
	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		time.Sleep(2 * time.Second)
		return true   //return false的时候就只会打印一组key，Range是接受到true的时候才会往下执行
	})

}

```
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	pool := &sync.Pool{}
	pool.Put(1)
	pool.Put(2)
	pool.Put(3)
	pool.Put(4)
	pool.Put(5)
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(pool.Get())
	}

}

```
#### pool里面没有了，就好循环打印nil
```go
2
1
3
5
4
<nil>
<nil>
<nil>
<nil>
<nil>


```


```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	codn := sync.NewCond(&sync.Mutex{})
	/*codn.L.Lock()
	codn.Wait()
	codn.Wait()
	codn.Wait()
	codn.Wait()
	codn.L.Unlock()
	codn.Signal()    //通知wait完成一个，可能有好多个wait
	codn.Broadcast() //广播*/
	go func() {
		codn.L.Lock()
		fmt.Println("lock1")
		codn.Wait()
		fmt.Println("unlock1")
		codn.L.Unlock()

	}()

	go func() {
		codn.L.Lock()
		fmt.Println("lock2")
		codn.Wait()
		fmt.Println("unlock2")
		codn.L.Unlock()
	}()

	time.Sleep(2 * time.Second)
	codn.Broadcast()
	fmt.Println("过1s之后通知解锁：")
	time.Sleep(1 * time.Second)
}

```
```go
lock2
lock1
unlock2
unlock1
过1s之后通知解锁：


```
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	codn := sync.NewCond(&sync.Mutex{})
	/*codn.L.Lock()
	codn.Wait()
	codn.Wait()
	codn.Wait()
	codn.Wait()
	codn.L.Unlock()
	codn.Signal()    //通知wait完成一个，可能有好多个wait
	codn.Broadcast() //广播*/
	go func() {
		codn.L.Lock()
		fmt.Println("lock1")
		codn.Wait()
		fmt.Println("unlock1")
		codn.L.Unlock()

	}()

	go func() {
		codn.L.Lock()
		fmt.Println("lock2")
		codn.Wait()
		fmt.Println("unlock2")
		codn.L.Unlock()
	}()

	time.Sleep(1 * time.Second)
	codn.Signal()
	time.Sleep(1 * time.Second)
	codn.Signal()
	time.Sleep(1 * time.Second)

}

```