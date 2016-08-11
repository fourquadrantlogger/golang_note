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
	m["1"]=40;
	n["3"]=60;
	fmt.Println(m,n)
	fmt.Println(m,n)
}
