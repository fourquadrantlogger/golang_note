package main

import "fmt"

func defer_call() {
	defer func() {
		fmt.Println("before recover")
	}()

	defer func() {
		e := recover()
		fmt.Println("recover ", e)
	}()

	defer func() {
		fmt.Println("after recover")
	}()

	panic("panic info")
}
func main() {
	defer_call()
}
