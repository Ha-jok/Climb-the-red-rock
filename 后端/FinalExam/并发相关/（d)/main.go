package main

import "fmt"

//使用并发操作可以提高程序的运行速度，请实现高并发求一百万之内的所有素数
//
//如果你参考了网络上的文章/程序，请附带上包含你个人理解的注释and被你参考文章/程序的链接


//分析，一百万太大，将其拆分成一千个进程。每次进程求不同范围（1000）内的素数

//创建一个函数，查找一千个数中的素数
func prime(start int,ch chan string){
	//设置一个起始点，从该起始点到以后一千个数字中查找素数。
	for st:=start;st<start+1000;st++{
		if st==0|1 {
			continue
		}
		//声明一个变量i，从2开始自增，
		//对st进行判断，若st为素数，则只有i=st时，才有st%i=0，可判
		var i int
		for i=2;i<=st;i++{
			if st%i==0 {
				break
			}
		}
		if i==st{
			fmt.Print(st,"为素数\t")
		}
		//将一个空字符传入通道
		ch <- ""
	}
}
//接收函数，用来接收通道中的值。
func receive(ch chan string){
	var receive string
	receive = <- ch
	fmt.Print(receive)


}
func main(){
	var ch = make(chan string,1000000)
	//起始点从0开始，每次自增1000，直到100000
	for n:=0;n<1000000;n=n+1000{
		go prime(n,ch)
	}
	//用一个for循环接收通道中的值
	for n:=0;n<999998;n++{
		receive(ch)
	}

}
