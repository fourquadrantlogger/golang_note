package main

import "fmt"

type A struct {
	Name string
}
type B struct {
	A    A
	Name string
}
type C struct {
	A    *A
	Name string
}

func main() {
	fmt.Println(new(C))
	fmt.Println(new(B))
}
