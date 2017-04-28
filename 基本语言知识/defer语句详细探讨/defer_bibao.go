package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func main() {
	a := 1
	b := 2

	defer calc("1", a, calc("10", a, b)) // "1",1,3("10",1,2,3),4
	a = 0
	defer calc("2", a, calc("20", a, b)) //"2",0,2("20",0,2,2),2
	b = 1
}
