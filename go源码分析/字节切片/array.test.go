package main

import "fmt"

type S struct {
}

func main() {
	len := 9
	a := [len]S{}
	fmt.Println(a)
}
