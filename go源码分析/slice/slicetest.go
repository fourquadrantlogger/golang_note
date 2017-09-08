package main

import (
	"fmt"

	"unsafe"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	p1 := unsafe.Pointer(&s[:0])
	//p3 :=  uintptr(s[:0])
	fmt.Printf(p1)
}
