package main

import "fmt"

func main() {
	var array  [5]float64=[5]float64{1,2,3,4,5}

	news:=array

	array[4]=5.1
	news[1]=1.1
	fmt.Println(array,news)
}
