package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {

		go func() {
			wg.Done()
			//defer wg.Add(-1)
			fmt.Println("i", i)
		}()
	}
	for i := 0; i < 10; i++ {
		go func(n int) {
			wg.Done()
			//defer wg.Add(-1)
			fmt.Println("j", n)
		}(i)
	}
	wg.Wait()
}
