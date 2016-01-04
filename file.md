#file相关处理包及字符串与二进制函数

##ioutil
ioutil非常简单，一次把文件的所有内容都读出来。如果文件较大，会占用很多内存，需要小心。
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
