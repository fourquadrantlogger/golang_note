package main

import "time"

func main() {

	print(1)
	a := make(chan int, 1)
	b := make(chan int, 1)
	b <- 1
	a <- 0
	for {
		select {
		case A := <-a:
			print(A)
			a <- 0
		case B := <-b:
			print(B)
			b <- 1
		}
		time.Sleep(time.Second)
	}

	print(2)
}
