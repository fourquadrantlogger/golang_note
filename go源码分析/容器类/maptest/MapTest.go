package main

import "fmt"

func main() {
	s:=make(map[string]int)
	s["1"]=0
	s["2"]=1

	value,have:=s["1"]
	fmt.Println(have,value)
	v,h:=s["3"]
	fmt.Println(h,v)
}
