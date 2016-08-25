package main

import (
	"fmt"
)

//假设A和B表示一个多线程的程序执行的两个操作。如果A happens-before B，那么A操作对内存的影响 将对执行B的线程(且执行B之前)可见。
//无论使用哪种编程语言，有一点是相同的：如果操作A和B在相同的线程中执行，并且A操作的声明在B之前，那么A happens-before B。
var c = make(chan int, 10)
var a string

func main(){
	go ff()
	<-c            // (3)
	fmt.Println(a) // (4)
}

/*
上述代码可以确保输出"hello, world"，因为(1) happens-before (2)，(4) happens-after (3)，再根据上面的第一条规则(2)是 happens-before (3)的，最后根据happens-before的可传递性，于是有(1) happens-before (4)，也就是a = "hello, world" happens-before print(a)。
*/
func ff() {
	a = "hello, world" // (1)
	c <- 0             // (2)
}