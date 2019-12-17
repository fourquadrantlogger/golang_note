package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	pase_student()
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
		//这里对&stu，变异后只会指向最后一个对象
	}
	bs, _ := json.Marshal(m)
	fmt.Println(string(bs))
	//{"li":{"Name":"wang","Age":22},"wang":{"Name":"wang","Age":22},"zhou":{"Name":"wang","Age":22}}
}
