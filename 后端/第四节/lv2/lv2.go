package main

//在每天凌晨两点，输出"谁能比我卷！"
//在每天早上六点，输出"早八算什么，早六才是吾辈应起之时"
//自程序运行时起，每过半分钟输出"芜湖！起飞！"
//其他你想完成的功能

//引入包
import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
//计时器1：每天凌晨两点，输出"谁能比我卷！"
func timer1() {
	t:=time.Now().Format("15:04:05")
	for  {
		if t=="02:00:00"{
			fmt.Println("谁能比我卷！")
		}
		time.Sleep(time.Second)
	}

}
//计时器2：每天早上六点，输出"早八算什么，早六才是吾辈应起之时"
func timer2(){
	t:=time.Now().Format("15:04:05")
	for  {
		if t=="06:00:00"{
			fmt.Println("八算什么，早六才是吾辈应起之时")
		}
		time.Sleep(time.Second)

	}
}
//计时器3：每过2秒钟输出"芜湖！ddz!"
func timer3(){
	t:=time.NewTicker(2*time.Second)
	for  {
		<- t.C
		fmt.Println("\"芜湖！ddz!")

	}
}

//主函数,只使用wg.Add和wg.Wait而不使用wg.Down使得程序一直运行。
func main(){
	wg.Add(1)
	go timer1()
	go timer2()
	go timer3()
	wg.Wait()
}