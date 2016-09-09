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

const (
	WatchState = 1 << iota
	MultiState
	SubscribeState
	MonitorState
)

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

	fmt.Println("WatchState", WatchState)
	fmt.Println("MultiState", MultiState)
	fmt.Println("SubscribeState", SubscribeState)
	fmt.Println("MonitorState", MonitorState)

}
