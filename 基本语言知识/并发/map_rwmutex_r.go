package main

import (
	"fmt"
	"sync"
)

var m map[string]interface{}

func main() {
	m = make(map[string]interface{}, 1024)
	l := sync.RWMutex{}
	go func() {
		//write
		i := 0
		for {
			l.RLock()
			m["key"] = (i)
			fmt.Print("w")
			i++
			l.RUnlock()
		}
	}()
	go func() {
		//read
		i := 0
		for {

			i += m["key"].(int)
			fmt.Print("r")

		}
	}()
	for true {

	}
}
