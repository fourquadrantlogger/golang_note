package main

import "fmt"

func main() {
	in:=make(chan int,1)

	in<-1

	i:=<-in
	fmt.Println(i)
	close(in)

	if v, ok:= <-in; ok {
		fmt.Println(v)
	}else {
		fmt.Println(ok)
	}
}
