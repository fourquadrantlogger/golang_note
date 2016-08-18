package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)



	go func() {

		time.Sleep(time.Second * 1)

		c1 <- "one"

	}()

	go func() {

		time.Sleep(time.Second * 2)

		c2 <- "two"

	}()



	for true {

		select {

		case msg1 := <-c1:

			fmt.Println("received", msg1)

		case msg2 := <-c2:

			fmt.Println("received", msg2)

		}

	}
}
