package main

import "fmt"

//基于管道，我们可以把打印的协程拓展为N个。
//
//请在main函数中开启10个协程输出一段话，要求10行话全部输出完毕后再结束main函数。

//声明两个函数，一个为go进程，一个为主进程
//go进程函数
func dio(num int,ch chan string){
	//声明一个字符串用于向通道中传输字符串
	var jojo string
	//进行for循环，使得每句话都不一样
	for i:=0;i<=num;i++{
		fmt.Print("木大")
		jojo =jojo+"欧拉"
	}
	fmt.Print("\n")
	//向通道中传输字符串
	ch <- jojo
}
//主进程中执行的函数，用于接收通道中的数值
func jojo(ch chan string)  {
	var jojo string
	jojo = <- ch
	fmt.Println(jojo)
}

func main (){
	//声明一个容量为9的通道，利用通道堵塞使得go进程执行完毕
	var ch = make(chan string,10)
	//开启10个go进程
	for num:=0;num<10;num++{
		go dio(num,ch)
	}
	//接收通道中的值
	for num:=0;num<10;num++{
		jojo(ch)
	}

}
//参考李文周博客并发：https://www.liwenzhou.com/posts/Go/14_concurrence/#autoid-1-3-3