package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

//不允许用户重复注册
//用户可以在登录之后修改自己的密码，并且程序退出后会用新的用户数据覆盖旧的用户数据
//使用crypto包中的加密算法将用户名密码加密后再写入到文件中，在读取时再进行解密。如果用到密钥，密钥可以自己输入也可以写在代码中。如果不知道使用哪种算法，建议使用hmac算法加密（答案将会使用这种算法）
//分析：由于空格也会被加密，所以先把数据加密再加上空格存入字符串，再将字符串存入文件
//功能1：如果注册已经注册过的用户名，提示用户名已存在
//功能2：登录后可选择修改密码
//功能3：加密（喵喵咪，渣渣猫需要查资料了）
//分析，需要一个加密函数，在储存到文件时进行加密。然后在用户名登录的时候对输入的字符进行加密。然后再进行匹配

//新增功能
//实现功能一:162到175，使用两次加密函数
//新增功能2：实现成功，114到125行，使用一次加密函数
//新增功能3：加密函数设置，29-34行
//功能三：要求，两个字符串，返回两个fu后的字符串。使用于存入文件和登录，登录修改密码
//第一次测试发现修改过的密码无法进行保存。解决思路。每一次都将字典重新转化为字符串保存.测试成功。代码实现于78到82行

//初始有：用户名ddz,密码yyds数据

//加密函数，对一个字符串进行加密，返回一个字符串.
func encry(s string)string{
	h:=md5.New()
	h.Write([]byte(s))
	x:=hex.EncodeToString(h.Sum(nil))
	return x
}







//定义一个结构体，一个字典，一个字符串
type data struct{
	username string `json:"Username,omitempty"`
	password string `json:"Password,omitempty"`
}
var datas=make(map[string]string,2)
var s string

//读取文件,并存入map未改动
func read (){
	//打开文件
	file,err:=os.Open("./src/prac/homework/第四节/lv3E/user.data")
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//读取文件
	var p=make([]byte,100000)
	n,err:=file.Read(p)
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//将文件读取到map
	var s1 string
	var ss=make([]string,n)
	s1=string(p[1:n-1])
	ss=strings.Fields(s1)
	for i:=0;i<len(ss);i=i+2{
		datas[ss[i]]=ss[i+1]
	}
	file.Close()
}

//写入文件分为增添和保存，为了实习修改密码函数，需要将保存函数独立
func save(){
	//提取字典信息
	s=""
	for un,pw:=range datas{
		s=s+" "+un+" "+pw
	}
	//将字符串json解码
	x,_:=json.Marshal(s)

	//以只写模式打开文件
	file,err:=os.OpenFile("./src/prac/homework/第四节/lv3E/user.data",os.O_WRONLY,0000)
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//将数据写入文件
	_,err=file.Write(x)
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	file.Close()
}
//将新的用户存入字符串，以便保存
func add (d data){
	s=s+" "+d.username+" "+d.password
}


//登录
func enter(un,pw string){
	//加密输入信息
	un=encry(un)
	pw=encry(pw)

	//判断登陆信息
	p:=datas[un]
	if p!="" {
		if p==pw {
			fmt.Println("登陆成功")
			//	实现修改密码功能
			fmt.Println("选择：\n1.修改密码\n2.退出登录")
			var sele int
			fmt.Scanln(&sele)
			if sele==1 {
				fmt.Print("请输入新密码：")
				var pw1 string
				fmt.Scanln(&pw1)
				pw1=encry(pw1)
				datas[un]=pw1
				strings.Replace(s,pw,pw1,1)
				save()
			}
		}
		if p!=pw{
			fmt.Println("您输入的密码有误，请重新输入")
		}
	}
	if p=="" {
		fmt.Println("您输入的用户名不存在，请注册")
	}
}

//注册，将新增用户存入结构体
var d data
func rigister(un,pw string) data {
	d.username=un
	d.password=pw
	datas[d.username]=d.password
	return d
}





func main(){
	//设置循环,以i作为条件，便于跳出循环
	i:=0
	read()
	//循环开始
	for i==0 {
		fmt.Println("1.登录\n2.注册\n3.退出程序")
		var in int
		var u,p string
		fmt.Scanln(&in)
		//选择不同功能
		switch in {
		case 1:{
			//实现登录
			fmt.Print("用户名：")
			fmt.Scanf("%s\n",&u)
			fmt.Print("密码：")
			fmt.Scanf("%s\n",&p)
			enter(u,p)
		}
		case 2:{
			//实现注册
			fmt.Print("用户名：")
			fmt.Scanf("%s\n",&u)
			//实现提示注册重复
			u=encry(u)
			if datas[u]!="" {
				fmt.Println("用户名重复,请更改您的用户名")
			}
			if datas[u]=="" {
				fmt.Print("密码：")
				fmt.Scanf("%s\n",&p)
				p=encry(p)
				//将新用户存入结构体
				d=rigister(u,p)
				//将新用户添加到数据中
				add(d)
				//保存信息
				save()
			}

		}
		case 3:
			i++
		}
	}
}