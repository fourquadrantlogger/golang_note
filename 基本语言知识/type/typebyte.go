package main

import "fmt"

type N byte

func (this *N) Decode() byte {
	return byte(*this)
}
func main() {
	n := N(1)
	fmt.Println(n.Decode())
}
