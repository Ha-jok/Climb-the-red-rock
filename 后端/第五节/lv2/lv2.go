package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

//分析需求，有注册，登录，修改密码，加密功能。
//需要有一个字典储存加密后的账号密码。并支持更改密码
//利用cypto包进行加密
//利用参数绑定进行登录。
//需要一个文件，储存信息，并读取信息。

//初始，账号：ddz,密码:yyds


//加密函数
func encry(s string)string{
	h:=md5.New()
	h.Write([]byte(s))
	x:=hex.EncodeToString(h.Sum(nil))
	return x
}

//读取文件，存入字典
var datas=make(map[string]string,100)
func read(){
	file,err:=os.Open("./homework/第五节/lv2/user.data")
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	//读取文件
	var p=make([]byte,10000)
	n,err:=file.Read(p)
	var s string
	var ss=make([]string,n)
	s=string(p[1:n-1])
	ss=strings.Fields(s)
	//将数据存入字典
	for i:=0;i<len(ss);i=i+2{
		datas[ss[i]]=ss[i+1]
	}
	file.Close()
}

//写入文件
func save(){
	//提取字典信息
	s:=""
	for un,pw:=range datas{
		s=s+" "+un+" "+pw
	}
	x,_:=json.Marshal(s)

	//只写方式打开文件
	file,err:=os.OpenFile("./homework/第五节/lv2/user.data",os.O_WRONLY,0000)
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}

	//写入数据
	_,err=file.Write(x)
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	file.Close()
}

//登录功能
var showstring string
func login(un,pw string){
	un=encry(un)
	pw=encry(pw)
	p:=datas[un]
	if p=="" {
		showstring="账号为空请注册"
	}else {
		if p!=pw {
			showstring="密码错误，请重新输入密码"
		}else {
			showstring="登录成功"
		}
	}
	fmt.Println(showstring)
}

//注册功能
func creatuser (un,pw string){
	un=encry(un)
	pw=encry(pw)
	if datas[un]!=""{
		showstring="账号重复，请重新注册"
	}else {
		datas[un]=pw
		showstring="注册成功"
		save()
	}

}

//修改密码功能
func editpw(un,pw string){
	un=encry(un)
	pw=encry(pw)
	datas[un]=pw
	save()
}

//创建cookie,来控制修改密码的权限
func cookie (c *gin.Context){
	value,err:=c.Cookie("cookie")
	if err!=nil{
		c.Abort()
	}else {
		c.Set("cookie",value)
	}
}

//主函数
func main(){
	read()
	//创建路由
	r:=gin.Default()
	//登录界面
	r.POST("/login",func (c *gin.Context){
		username:=c.PostForm("username")
		password:=c.PostForm("password")
		login(username,password)
		if showstring=="登录成功" {
			c.SetCookie("cookie",username,600,"/","",false,true)
			c.String(200,showstring)
		}
	})
	//修改密码界面
	r.POST("/edit",cookie,func (c *gin.Context){
		username1,_:=c.Get("cookie")
		username:=username1.(string)
		password:=c.PostForm("password")
		editpw(username,password)
		c.JSON(200,gin.H{
			"username":username,
			"password":password,
		})
	})
	//注册界面
	r.POST("/creatuser",func (c *gin.Context){
		username:=c.PostForm("username")
		password:=c.PostForm("password")
		creatuser(username,password)
		c.String(200,showstring)
	})
	//查看界面
	r.GET("/hello",cookie,func(c *gin.Context){
		cookie,_:=c.Get("cookie")
		str:=cookie.(string)
		c.String(200,str+"期待ddz女装")
	})
	r.Run()
}