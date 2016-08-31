package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	slice := arr[1:2]
	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println(len(slice))
	slice = append(slice, 6, 7, 8)
	fmt.Println(slice)
	fmt.Println(arr)

	//[2]
	//4
	//1
	//[2 6 7 8]
	//[1 2 6 7 8]
}
