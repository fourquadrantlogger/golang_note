package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	js_str := []byte(`{"age":1}`)
	var value map[string]interface{}

	json.Unmarshal(js_str, &value)
	fmt.Println(reflect.TypeOf(value["age"]))
	age := value["age"]

	fmt.Println(value)
	fmt.Println(reflect.TypeOf(age))
	fmt.Println(reflect.TypeOf(1))
}

///NND

// float64
// map[age:1]
// float64
// int
// 默认居然是float64类型
