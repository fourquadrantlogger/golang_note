# 匿名结构体

```
package main
     
import (
    "fmt"
)
     
func main() {
    var user struct{Name string; Gender int}
    user.Name = "dotcoo"
    user.Gender = 1
    fmt.Printf("%#v\n", user)
}
```