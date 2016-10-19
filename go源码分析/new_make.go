package main

import "fmt"

type d struct {
	D string
}

func main() {
	d := new(d)
	fmt.Println(d)
}
