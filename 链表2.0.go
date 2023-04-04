package main

import "fmt"

type NodeData struct {
	data interface{}
	next *NodeData
}

type Stack02 struct {
	top  *NodeData
	size int
}

func (stack *Stack02) Push(val interface{}) {
	//初始化赋值一个新的指针变量,指向新入栈的节点
	newNode := &NodeData{val, nil}
	if stack.top == nil {
		stack.top = newNode
	} else { //栈不为空
		newNode.next = stack.top
		stack.top = newNode
	}
	stack.size++
}

func (stuck *Stack02) Pop() interface{} {
	if stuck.top == nil {
		return nil
	} else { //栈不空
		//栈顶是链表的头部
		val := stuck.top.data
		//删除节点
		stuck.top = stuck.top.next
		stuck.size--
		return val
	}
}

func main() {
	s := &Stack02{}
	s.Push("你好")
	s.Push(888)
	s.Push('a')
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
