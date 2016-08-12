package main

import (
	"sort"
	"fmt"
)

func main()  {
	slice_m:=[]int{3,2,4,5,3,5,25,3,2,1,4,2,4,6,63,2,3,3}
	fmt.Println(slice_m)
	sort.Ints(slice_m)
	fmt.Println(slice_m)
}