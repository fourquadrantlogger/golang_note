package main

import "fmt"

func main() {
	v := f1()
	fmt.Println(v)
}

func f1() (i int) {

	defer func() {
		i++
	}()
	return 0
}

func f2() (i int) {

	defer func(i int) {
		i++
	}(i)
	return 0
}
