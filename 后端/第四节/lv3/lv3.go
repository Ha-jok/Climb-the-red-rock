package main

//程序启动时，从本地文件（文件名随意，建议使用users.data）读入用户名和密码并保存到map中，规则是用户名不能重复，密码长度大于6位（可自行增加其他规则）。
//用户注册时，判断用户名密码是否符合规范并输出提示信息
//用户选择退出程序（不是人为强行终止程序）时，将新注册的用户信息也写入到文件中，然后退出程序

//需要1：注册，保存，登录功能。2，读取文件，写入文件

//引入包
import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)
//定义一个结构体和一个字典
type data struct{
	username string `json:"Username,omitempty"`
	password string `json:"Password,omitempty"`
}
var datas=make(map[string]string,2)
var s string
//读取文件,并存入map
func read (){
	file,err:=os.Open("./src/prac/homework/第四节/lv3/user.data")
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var p=make([]byte,100000)
	n,err:=file.Read(p)
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var ss=make([]string,n)
	s=string(p[1:n-1])
	ss=strings.Fields(s)
	for i:=0;i<len(ss);i=i+2{
		datas[ss[i]]=ss[i+1]
	}
	file.Close()
}
//写入文件
func write (d data){

	s=s+" "+string(d.username)+" "+string(d.password)
	fmt.Println(s)
	x,_:=json.Marshal(s)
	fmt.Println(string(x))
	file,err:=os.OpenFile("./src/prac/homework/第四节/lv3/user.data",os.O_WRONLY,0000)
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_,err=file.Write(x)
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	file.Close()
}


//登录
func enter(un,pw string){
	fmt.Println(un)
	fmt.Println(datas[un])
	p:=datas[un]
	if p!="" {
		if p==pw {
			fmt.Println("登陆成功")
		}
		if p!=pw{
			fmt.Println("您输入的密码有误，请重新输入")
		}
	}
	if p=="" {
		fmt.Println("您输入的用户名不存在，请注册")
	}
}

//注册
var d data
func rigister(un,pw string) data {
	d.username=un
	d.password=pw
	datas[d.username]=d.password
	return d
}





func main(){
	i:=0
	read()
	for i==0 {
		fmt.Println("1.登录\n2.注册\n3.退出程序")
		var in int
		var u,p string
		fmt.Scanln(&in)
		switch in {
		case 1:{
			fmt.Print("用户名：")
			fmt.Scanf("%s\n",&u)
			fmt.Print("密码：")
			fmt.Scanf("%s\n",&p)
			enter(u,p)
		}
		case 2:{
			fmt.Print("用户名：")
			fmt.Scanf("%s\n",&u)
			fmt.Print("密码：")
			fmt.Scanf("%s\n",&p)
			d=rigister(u,p)
			//fmt.Println(d)测试使用
			write(d)
		}
		case 3:
			i++
	}
  }
}