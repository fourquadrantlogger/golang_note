package main

import "fmt"

func main() {
	bucketCntBits := 3
	bucketCnt := 1 << bucketCntBits
	fmt.Println(bucketCnt)
}
