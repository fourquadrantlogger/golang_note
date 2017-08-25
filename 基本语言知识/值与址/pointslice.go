package main

import "fmt"

type A struct {
	B int
}

func main() {
	l := make([]*A, 0)
	in := make([]int, 0)
	for _, a := range l {
		in = append(in, a.B)
		fmt.Println(a, a.B, in)
	}

}
