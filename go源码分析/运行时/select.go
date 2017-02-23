package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	for {
		int_chan := make(chan int, 1)
		string_chan := make(chan string, 1)
		int_chan <- 1
		string_chan <- "呵呵"

		select {
		case v := <-int_chan:
			fmt.Println(v, "out put")
		case v := <-string_chan:
			panic(v + "out put")
		}
	}
}

// 此题目，说明select函数内部的case，是随机的
