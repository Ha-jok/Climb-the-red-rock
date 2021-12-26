package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	go func() {
		fmt.Println("有点强人锁男")
		mu.Lock()
	}()

	mu.Unlock()
}
//该代码将加锁代码放到了go进程中。
//
//加锁应该在go进程之外（前）进行。

