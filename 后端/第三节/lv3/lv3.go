package main

//修改代码，输出0-9

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	over := make(chan bool ,1)  //创建布尔类型通道，为零缓冲，会造成堵塞，根据代码，该通道只会接收一个变量，所以设容量为1
	for i := 0; i < 10; i++ {
		wg.Add(1)//一个十次循环体
		go func(x int) {
			//相当于构建了十个goroutine进程。
			//错误发现，go进程是同时进行的，i是会变化的，即i=1建立的go进程在输出时，i的值未知
			//思路直接将i作为参数传入
			fmt.Println(x)
			wg.Done()
		}(i)
		if i == 9 {
			over <- true
			_=<-over
		}
		wg.Wait()
	}

	fmt.Println("over!!!")
}