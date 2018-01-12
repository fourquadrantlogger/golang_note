package main

import (
	"strconv"
	"math"
	"fmt"
)

type info struct {
	 province string
	city string
	lat float64
	lng float64
}
func main() {
	m:=make(map[int64]info)
	for i :=0;i<1000*1000*10;i++{
		m[int64(i)]=info{
			strconv.Itoa(i),
			strconv.Itoa(i),
			math.Sin(float64(i)),
			math.Cos(float64(i)),
		}
	}
	print("ok")
	for{
		fmt.Scan()
	}
}
