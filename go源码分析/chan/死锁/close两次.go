package main

import "fmt"

func main() {
	c := make(chan int)
	close(c)

	i, isClose := <-c
	fmt.Println(i, isClose)
}
