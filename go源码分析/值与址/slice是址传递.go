package main

import "fmt"

func main() {
	var slice []float64=[]float64{1,2,3,4,5}

	news:=slice

	slice[4]=5.1
	news[1]=1.1
	fmt.Println(slice,news)
}
