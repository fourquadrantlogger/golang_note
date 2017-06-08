package main

import "fmt"

func remoteslice(s *([]byte)) {
	s = &([]byte{0x1, 0x2})
}
func remotemap(m map[string]string) {
	m["s"] = "sdf"
}
func main() {
	s := []byte{0x1, 0x2, 3, 4, 5, 6, 7, 8, 9, 0}
	remoteslice(&s)
	fmt.Println(s)

	m := make(map[string]string)
	remotemap(m)

	fmt.Println(m)
}
