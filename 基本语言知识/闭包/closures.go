package main

import "fmt"

func main() {
	var closures [2]func()
	for i := 0; i < 2; i++ {
		closures[i] = func() {
			fmt.Println(i)
		}
	}
	closures[0]()
	closures[1]()
}
