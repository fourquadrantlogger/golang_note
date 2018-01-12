// 会导致并发冲突的map访问
// fatal error: concurrent map writes
package main

import (
	"fmt"
	"strconv"
)

func main() {
	m := make(map[int]string, 1024)
	var closures [3]func(int)
	for i := 0; i < 3; i++ {
		closures[i] = func(v int) {
			for {
				m[v] += strconv.Itoa(v)
				if len(m[v])%1000 == 0 {
					fmt.Println(v, len(m[v]))
				}
			}
		}
	}
	go closures[0](0)
	go closures[1](1)
	go closures[2](2)
	for {
		fmt.Scanln()
	}
}
