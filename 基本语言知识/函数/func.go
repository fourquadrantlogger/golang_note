package main

import "fmt"

type Atype struct {
	name string
}

func do(this *Atype) (result *Atype) {
	this.name = "lipeng"
	return this
}

func doing() (result Atype) {
	return
}
func domap() (result map[string]interface{}) {
	return
}
func main() {
	var a Atype
	fmt.Println(do(&a))
	fmt.Println(domap())
}
