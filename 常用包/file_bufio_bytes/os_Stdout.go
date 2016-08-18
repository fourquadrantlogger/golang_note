package main

import (
	"os"

	"fmt"
	"io"

)

func main() {
	w:=os.Stdout
	w.WriteString("hello")


	f,err:=os.Create("wt.txt")
	if err!=nil{
		panic(err)
	}
	written,err:=io.Copy(f,w)
	if err!=nil{
		panic(err)
	}
	fmt.Println(written)
}
