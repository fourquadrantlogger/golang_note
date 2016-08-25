package main

import "fmt"

func main() {
	var s="我是字符串"
	fmt.Println(&s)
	var s2=s
	fmt.Println(&s2)


	var s3 ="我"
	s_string(s3)
	fmt.Println(&s3,len(s3))
	s3+="1"
	s_string(s3)
	fmt.Println(&s3,len(s3))
	s3+="2"
	s_string(s3)
	fmt.Println(&s3,len(s3))
}
func s_string(s string)  {
	fmt.Println("func",&s,"发生了值copy")

}
