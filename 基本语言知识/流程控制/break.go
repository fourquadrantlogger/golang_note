package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			switch i * j % 10 {
			case 1:
				fmt.Println(i, j)
				break

			}
			fmt.Println("break 0")
		}
		fmt.Println("break 1")
	}
	fmt.Println("break 2")
}
