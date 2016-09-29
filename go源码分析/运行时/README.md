# https://golang.org/pkg/runtime/

## 常用函数

+ func GOROOT() string

+ func NumCPU() int
NumCPU returns the number of logical CPUs usable by the current process.

+ func NumGoroutine() int
NumGoroutine returns the number of goroutines that currently exist.

+ func Version() string
, a release tag like "go1.3"

+ func (f *Func) Name() string
Name returns the name of the function.