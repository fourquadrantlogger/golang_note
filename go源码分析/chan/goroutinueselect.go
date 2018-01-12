package main

import (
	"fmt"
	"time"
)

func gorouting(ch chan int) {

	i := 0

	i++
	for {
		ch <- i
		fmt.Println(i)
		<-ch
		i++
		time.Sleep(time.Second)
	}
}
func main() {
	ch := make(chan int, 1)
	go gorouting(ch)
	for {
		ch <- 0
		fmt.Println("I want runing")
		<-ch
		time.Sleep(time.Second * 5)
	}

}
