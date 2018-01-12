package main

import (
	"fmt"
	"github.com/timeloveboy/moegraphdb/graphdb"
)

func main() {
	//ints:=make([]int,8*10000*10000)
	// 197mb
	//ints:=make([]int64,8*10000*10000)
	// 197mb
	//ints:=make(map[uint64]uint64,8*10000*10000)
	// 594mb
	//ints:=make([]graphdb.User,8*10000*10000)
	// 590mb
	ints := make(map[uint]*graphdb.User)
	// 594mb
	for k := uint(0); k < 4*1000*10000; k++ {
		ints[k] = new(graphdb.User)
		ints[k].Uid = k
		ints[k].Fans = make(map[uint]byte, 0)
		ints[k].Likes = make(map[uint]byte, 0)
	}
	fmt.Println(len(ints))
	for true {

	}
}
