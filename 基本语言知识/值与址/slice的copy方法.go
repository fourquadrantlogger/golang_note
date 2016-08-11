package main

import "fmt"
/**
该函数主要是切片（slice）的拷贝，不支持数组
将第二个slice里的元素拷贝到第一个slice里，拷贝的长度为两个slice中长度较小的长度值
示例：
 */
func main() {
	s := []int{1,2,3}
	fmt.Println(s) //[1 2 3]
	copy(s,[]int{4,5,6,7,8,9})
	fmt.Println(s) //[4 5 6]
}
