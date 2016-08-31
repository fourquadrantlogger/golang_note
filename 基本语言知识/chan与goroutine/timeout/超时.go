package main

import (
	"fmt"
	"time"
)


func main() {
	done := make(chan int, 1)
	go func() {
		time.Sleep(4 * time.Second)
		done <-0
	}()

	select {
	case result := <-done:
	fmt.Println(result)
	case <- time.After(3 * time.Second):
	fmt.Println("timeout")
	}
}
