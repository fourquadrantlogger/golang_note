package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		defer func(v int) {
			fmt.Println("num=", v)
		}(i)
	}
}

//
