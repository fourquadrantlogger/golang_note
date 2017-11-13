package main

import (
	"fmt"
	"sync"
	"time"
)

type Set struct {
	sync.RWMutex
	S map[interface{}]struct{}
}

func (set *Set) Iter2() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer func() {
			r := recover()
			if r != nil {
				fmt.Println("捕获bug")
			}
			set.RUnlock()
		}()

		set.RLock()
		for k := range set.S {
			ch <- k
		}
		close(ch)
	}()
	return ch
}

func main() {
	s := Set{}
	s.S = make(map[interface{}]struct{})
	s.S[1] = struct{}{}
	s.S[2] = struct{}{}

	for k := range s.Iter2() {
		fmt.Println("---", k)
		panic("测试bug")
	}

	time.Sleep(3 * time.Second)
}
