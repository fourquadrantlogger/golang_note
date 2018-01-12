package main

import (
	"fmt"
	"reflect"
)

type MoeInt int

func (this *MoeInt) Say() {
	fmt.Println(*this)
}
func (this *MoeInt) say() {
	fmt.Println(*this)
}
func main() {
	//反射类型
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))

	//反射函数
	var m MoeInt = 4
	i := &m
	object := reflect.ValueOf(i)
	object.MethodByName("Say").Call(nil)
	object.MethodByName("say").Call(nil)

}
