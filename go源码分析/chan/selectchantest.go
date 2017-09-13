package main

import "fmt"

var chan3 = make(chan int, 10)
var chan4 = make(chan int, 10)
var chs = []chan int{chan3, chan4}
var numbers = []int{1, 2, 3, 4, 5}

func main() {
	select {
	case getChan(0) <- getNumber(2):
		fmt.Printf("1st case is selected!\n")
	case getChan(1) <- getNumber(3):
		fmt.Printf("2st case is selected!\n")
	default:
		fmt.Printf("default!\n")
	}
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]

}

func getChan(i int) chan int {
	fmt.Printf("chs[%d]\n", i)
	return chs[i]

}
