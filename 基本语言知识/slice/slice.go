package main

import "fmt"

func main() {
	s:=make([]int,5)
	s=append(s,0)
	s=append(s,1)
	s=append(s,2)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	//8
	//10

	s2:=make([]int,0)
	fmt.Println(cap(s2))
	s2=append(s2,0)
	s2=append(s2,0)
	s2=append(s2,0)
	s2=append(s2,0)
	s2=append(s2,0)
	fmt.Println(cap(s2))
	//0
	//8

	s3:=make([]int,100)
	s3=append(s3,1)
	fmt.Println(cap(s3))
	fmt.Println(len(s3))
	//208
	//101

}

// 切片
type Slice struct {
  Data *interface{}
  Len int
  Cap int
}

