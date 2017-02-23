package main

import "fmt"

type Parent struct {
}

func (p *Parent) MethodB() {
	fmt.Println("MethodB from parent")
}

func (p *Parent) MethodA() {
	fmt.Println("MethodA from parent")
	p.MethodB()
}

type Child struct {
	Parent
}

func main() {
	c := Child{}
	c.MethodA()
}

//这道题目，
