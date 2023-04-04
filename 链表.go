package main

import "fmt"

type Node struct {
	data interface{}
	next *Node //node是指针变量
}

//在链表实现的栈中，栈底是链表的尾部，栈顶则是链表的头部
type Stack01 struct {
	//代表链表的头部
	top  *Node //top是一个指针变量，存的是地址，是Node这个元素类型的地址
	size int
}

/*
在 Go 语言中，方法 (Method) 是一种特殊的函数，它在定义时在其名称之前附加了一个接收器 (Receiver)，
这个接收器相当于该方法的调用者。而 (s *Stack) 就是一个接收器，表示 Push 方法属于 Stack 结构体类型，并且该方法操作 Stack 类型的对象。
具体地说，这里的 (s *Stack) 表示将 Stack 结构体类型的指针作为方法的接收器。
这样，调用该方法时，就可以通过 s 操作 Stack 类型的对象，例如可以访问 Stack 中的成员变量或调用其他方法。*/

//可以用s来操作Stuck类型的对象，该方法的调用者，这个方法属于Stuck这个结构体类型
func (s *Stack01) Push(val interface{}) {
	//newNode 是一个指向 Node 结构体的指针，&取地址符号，newNode存储的是 Node 结构体的地址
	newNode := &Node{val, nil} //&Node{}初始化一个Node类型的实例,
	if s.top == nil {
		//说明栈空，NewNode直接入栈
		s.top = newNode
	} else {
		//栈非空
		//newNode = stack.top,这样导致与新入栈的节点的断开了
		/*
			该结构体包含两个属性：
				data：存储节点中的数据；
				next：存储指向下一个节点的指针。*/
		newNode.next = s.top
		s.top = newNode
	}
	s.size++
}

//严重的错误，s Stack01这里调用的是值传递，操作的是这个栈的副本，而不是对原始的指向Stack01操作的
//可以将 Pop() 方法的接收器改为指向结构体的指针类型
func (s *Stack01) Pop() interface{} {
	if s.top == nil {
		return nil
	} else {
		//移除链表的头部
		value := s.top.data
		s.top = s.top.next
		s.size--
		return value
	}
}

func main() {
	stack := &Stack01{} //Push和Pop方法要用指向Stack实例的指针，如果写成stack := Stack01{}，只是创建了一个Stack空的实例而已，而不是一个指向实例的指针
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
