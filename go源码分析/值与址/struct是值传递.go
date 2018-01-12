package main

import "fmt"

func main() {
	n:=N{255,255,255}
	m:=n
	n.R=0
	m.B=0
	fmt.Println(n,m)

	p:=&N{255,255,255}
	q:=p
	p.R=0
	q.B=0
	fmt.Println(p,q)
}

type N struct {
	R,G,B int
}
