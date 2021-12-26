package main

import "fmt"

//func main() {
//
//	go func() {
//		fmt.Println("下山的路又堵起了")
//	}()
//
//}

//修改
func main() {
	//创建一个无缓冲通道，借助通道堵塞，完成并发进程完成
	var ch = make(chan string)
	go func() {
		//将字符串传输到通道中，等待主协程接收。
		ch <- "OK"
		fmt.Println("下山的路又堵起了")
	}()
	//接收通道中的数值，和go并发中对应，
	receive := <- ch
	//打印从通道中接收的值，即只有go进程中传输，receive才可以接收，达到了go进程执行后主进程才结束
	fmt.Println(receive)
}
//参考：红岩后端第三节课课件：https://www.yuque.com/gyxffu/uv3zph/ft77rx#0190f8d5
