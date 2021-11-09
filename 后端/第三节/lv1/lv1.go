package main
//要求：使用两个goroutine轮流打印100以内的数
//轮流打印，两个gotouyine应该是串行形式，思路：互斥锁
import (
	"fmt"
	"sync"
)

//声明变量，一个WaitGroup,一个互斥锁
var wg sync.WaitGroup
var lock sync.Mutex

//打印函数
func prin ()  {
	lock.Lock()
	for i:=1;i<=100;i++ {
		fmt.Println(i)
	}
	wg.Done()
	lock.Unlock()
}

func main (){
	wg.Add(2)
	go prin()
	go prin()
	wg.Wait()
}