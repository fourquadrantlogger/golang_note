package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 生成队列
	l := list.New()
	// 入队, 压栈
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	// 出队
	i1 := l.Front()
	l.Remove(i1)
	fmt.Printf("%d\n", i1.Value)
	// 出栈
	i4 := l.Back()
	l.Remove(i4)
	fmt.Printf("%d\n", i4.Value)
}
