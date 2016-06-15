# go 时间类型
time类型的String()方法，会生成

```go
2006-01-02 15:04:05.999999999 -0700 MST
```
从字符串截取前19个字符，刚好就是我们需要的
```go
import (
	"time"
	)

StringHelper.Substr(time.Now().String(),0,19)

2006-01-02 15:04:05
```
