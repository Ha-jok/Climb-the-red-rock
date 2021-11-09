package main
//a.使用os.Create函数，在你的项目目录下创建一个"plan.txt"文件，

//b.使用file.Write将"I’m not afraid of difficulties and insist on learning programming",写入"plan.txt"中。

//c.使用file.Read函数读取"plan.txt"的内容，并打印到控制台

import (
	"fmt"
	"os"
)



//函数a，创建文件
func creat ()  {
	file,err:=os.Create("./src/prac/homework/第三节/lv2/plan.txt")
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

}

//写入字符串
func wri (){
	file,err:=os.OpenFile("./src/prac/homework/第三节/lv2/plan.txt",os.O_WRONLY,0666)
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_,err=file.Write([]byte("I’m not afraid of difficulties and insist on learning programming"))
	fmt.Println(len("I’m not afraid of difficulties and insist on learning programming"))
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//读取函数
func read () {
	file,err:=os.Open("./src/prac/homework/第三节/lv2/plan.txt")
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	 p:=make([]byte,67)
	n,err:=file.Read(p)
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(p[:n]))
}


func main () {
	creat()
	wri()
	read()
}