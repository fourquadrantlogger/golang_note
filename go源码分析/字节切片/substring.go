package main

import "fmt"

func main() {
	a:="中12国123abc的的的"
	pre:="中12国"

	fmt.Println(substring(a,pre))
}

func substring(str,prefix string)( string){
	result:=[]rune(str)[len([]rune(prefix)):]
	return string(result)
}