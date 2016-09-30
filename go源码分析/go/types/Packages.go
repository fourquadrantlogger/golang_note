package main

import (
	"fmt"
	"go/types"
)

func main() {
	p := types.NewPackage("github.com/golangframework", "map256")

	p.MarkComplete()
	fmt.Println(p, p.Complete())
}
