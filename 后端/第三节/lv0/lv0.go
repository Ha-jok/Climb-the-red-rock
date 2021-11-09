package main
//要求：利用channel替代加锁
//分析add函数可知，x返回值为初值+50000
//main函数输出为100000
//思路：使用一个通道。利用管道堵塞，使两个goroutine串行执行

import (
	"fmt"
	"sync"
)

var x int
var wg sync.WaitGroup
//var mu sync.Mutex
var ch=make(chan int,1)
//定义一个int通道

func add(ch chan int) {
	for i := 0; i < 50000; i++ {
		ch <- x
		x = x + 1
		_=<-ch
	}
	wg.Done()
}



func main() {
	wg.Add(2)
	go add(ch)
	go add(ch)
	//连续进行两个goroutine进程，访问同一个函数。
	wg.Wait()
	fmt.Println(x)
}
