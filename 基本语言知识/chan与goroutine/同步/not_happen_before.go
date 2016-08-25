package main

import (
	"fmt"
)

var c = make(chan int, 1)
var a string

func f() {
	a = "hello, world" // (1)
	<-c                // (2)
}
func main() {
	go f()
	c <- 0         // (3)
	fmt.Println(a) // (4)
}

/*
这段代码不再有任何同步保证，使得(2) happens-before (3)
*/
