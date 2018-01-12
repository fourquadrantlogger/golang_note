package main

import "fmt"

func main() {
	var m map[string]int=map[string]int{
		"1":10,
		"2":20,
		"3":30,
	}

	var n map[string]int=map[string]int{
		"1":10,
		"2":20,
		"3":30,
	}
	n=m
	changemap(m)
	changemap(n)
	fmt.Println(m,n)
	fmt.Println(m,n)
}
func changemap(m map[string]int)  {
	m["1"]=40;
	m["3"]=60;
}