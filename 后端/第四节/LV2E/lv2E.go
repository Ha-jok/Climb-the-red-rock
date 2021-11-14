package main

//用户可以自己增加定时器（一次性的&重复提醒的）
//用户可以指定删除某个已存在的定时器
//用户可以在不删除定时器（重复提醒的）的情况下，取消定时器的下一次提醒

//分析：
//第一反应：我的脑壳疼欸
//这是直接要搞出来个闹钟程序欸
//功能：1，添加一次性计时器，2，添加重复计时器，3，删除计时器，4.取消下一次计时器提醒，增添功能：查看当前计时器
//思路，不能使用函数，进一步思考，决定选择字典储存计时器键为时间点，值为输出内容

//引入包
import (
	"fmt"
	"time"
)
//声明字典，容量设为20
var alarmClock=make(map[string]string,20)
var alarmClockOnce=make(map[string]string,20)

//实现功能一
func once(){
	fmt.Print("15:04（参考格式）\n请输入时间：")
	var timer string
	fmt.Scanln(&timer)
	fmt.Print("请输入提醒信息：")
	var timerValue string
	fmt.Scanln(&timerValue)
	alarmClockOnce[timer]=timerValue
}
//实现功能2
func timers(){
	fmt.Print("15:04（参考格式）\n请输入时间：")
	var timer string
	fmt.Scanln(&timer)
	fmt.Print("请输入提醒信息：")
	var timerValue string
	fmt.Scanln(&timerValue)
	alarmClock[timer]=timerValue
}
//实现功能3
func deleteTimer(){
	fmt.Println("选择删除计时器的类型：\n1.一次性计时器\n2.重复性计时器\n请注意删除已经存在的计时器")
	var sele int
	fmt.Scanln(&sele)
	fmt.Println("请输入要删除的计时器时间")
	var sele1 string
	fmt.Scanln(&sele1)
	switch sele {
	case 1:
		delete(alarmClockOnce,sele1)
	case 2:
		delete(alarmClock,sele1)
	}
}
//实现增添功能
func show(){
	x:=1
	for i,j:=range alarmClock{
		fmt.Println("重复性计时器",x,":",i,":",j)
		x++
	}
	for i,j:=range alarmClockOnce{
		fmt.Println("一次性性计时器",x,":",i,":",j)
		x++
	}
}


//主体函数
func main (){
	//测试部分
	alarmClock["22:37"]="阿巴阿巴"
	alarmClock["16:56"]="adf"
	alarmClockOnce["15:35"]="asdf"

	//主体部分
	for  {
		//判断部分
		//获取现在时间
		t:=time.Now().Format("15:04")
		//当现在时间属于字典的键时，即计时器到达预设时间点，执行计时器
		if alarmClock[t]!="" {
			fmt.Println(alarmClock[t])
		}
		if alarmClockOnce[t]!="" {
			fmt.Println(alarmClockOnce[t])
			delete(alarmClockOnce,t)
			//	一次性计时器执行一次后，删除该计时器
		}

		//交互部分
		fmt.Println("1.添加一次性计时器\n2.添加重复计时器\n3.删除计时器\n4.查看计时器")
		var sele int
		fmt.Scanln(&sele)
		switch sele {
		case 1:
			once()
		case 2:
			timers()
		case 3:
			deleteTimer()
		case 4:
			show()
		}
		time.Sleep(time.Second)
	}


}
