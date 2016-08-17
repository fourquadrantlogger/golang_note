package main

import "fmt"

func main() {
	c := make(chan int)  // Allocate a channel.

	// Start the sort in a goroutine; when it completes, signal on the channel.

	go func() {

		fmt.Println("c")

		c <- 1  // Send a signal; value does not matter.

	}()

	doSomethingForAWhile()

	<-c   // Wait for sort to finish; discard sent value.
}
