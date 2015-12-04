# string类型与rune

##len()
```go
s:="go语言编程"
var l int
l=len([]rune(s))
```

##substring()
```go
s:="go语言编程"
rs:=[]rune(s)
var str=rs[2:6]
fmt.Println(string(str))

```