# WaitGroup

先说说WaitGroup的用途：它能够一直等到所有的goroutine执行完成，并且阻塞主线程的执行，直到所有的goroutine执行完成。

Add(delta int)

Done()

Wait()

```
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
    wg.Add(20)
	for i := 0; i < 10; i = i + 1 {
		
		go func() {
			//wg.Done()
			//defer wg.Add(-1)
			EchoNumber(i)
		}()
	}
	for i := 0; i < 10; i = i + 1 {
		
		go func(n int) {
			//wg.Done()
			//defer wg.Add(-1)
			EchoNumber(n)
		}(i)
	}
	
	wg.Wait()
}

func EchoNumber(i int) {
	time.Sleep(3e9)
	fmt.Println(i)
}
```