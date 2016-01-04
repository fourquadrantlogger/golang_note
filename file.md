#file相关处理包及字符串与二进制函数

ioutil
``` go
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	buff, err := ioutil.ReadFile("d:\\test.txt")
	if err != nil {
		panic("open file failed!")
	}
	fmt.Print(string(buff))
}
```
