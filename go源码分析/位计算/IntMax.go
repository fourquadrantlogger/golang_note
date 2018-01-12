package main

import "fmt"

func main() {
	fmt.Println(int64(-1 << 63))
	fmt.Println(int64(1<<63 - 1))
}
