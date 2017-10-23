package main

import (
	"fmt"
	"time"
)

func main() {
	//时间戳
	t := time.Now()

	fmt.Println(t.Weekday().String())
	fmt.Println(int(t.Weekday()))
	fmt.Println(int(t.Month()))
	fmt.Println(int(t.Day()))
	y, m, d := t.Date()
	fmt.Println(y, m, d)
}
