package main

import (
	"fmt"
	"time"
)

func main() {
	d := make(map[int]int, 2000*10000)
	t1 := time.Now().Unix()
	for i := 0; i < 2000*10000; i++ {
		d[i] = i
	}
	fmt.Println(len(d))
	fmt.Println(time.Now().Unix() - t1)
}
