package main

import "fmt"

var days = [...]string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

func main() {
	var a interface{}
	a = days
	fmt.Println(a)
	switch vtype := a.(type) {
	default:
		fmt.Println(vtype)
	}
	fmt.Println()
	fmt.Println(len(days))

}
