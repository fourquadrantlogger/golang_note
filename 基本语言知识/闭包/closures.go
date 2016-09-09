package main

import (
	"fmt"
	"time"
)

type colsure func(*int)

var add colsure = func(v *int) {
	for {
		(*v)++
	}

}
var miux colsure = func(v *int) {
	for {
		(*v)--
	}
}

func main() {
	i := 0
	go add(&i)
	go miux(&i)

	for {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
