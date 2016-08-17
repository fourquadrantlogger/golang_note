package main

import "fmt"

func main() {
	fmt.Println("1")
	c := make(chan int)  // Allocate a channel.
	// Start the sort in a goroutine; when it completes, signal on the channel.
	go func() {
		fmt.Println("2")
		c <- 1  // Send a signal; value does not matter.
		fmt.Println("3")
	}()
	fmt.Println("4")
	<-c   // 主协程 等待
	fmt.Println("5")
}
