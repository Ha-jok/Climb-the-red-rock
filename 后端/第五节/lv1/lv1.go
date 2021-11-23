package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//分析要求，实现一个cookie，实现登录状态，并增加权限
//借助课件代码思路
//首先，建一个路由，发送账户密码，如果正确，将数据保存在服务端，实现cookie
//利用cookie的有无控制权限。


//定义登录中间件
func login (c *gin.Context){
	//向服务器端发送账号密码
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	//若账号密码正确，保存cookie
	if username=="ddz"&&password=="yyds" {
		//向客户端发送cookie
		c.SetCookie("cookie",username,600,"/","",false,true)
		c.JSON(200,gin.H{
			"message":"登陆成功",
		})
	}else {
		c.JSON(200,gin.H{
			"message":"登录失败，账号密码错误",
		})
	}

}


//实现cookie（函数形式）
func cookie (c *gin.Context){
	//获取cookie
	value,err:=c.Cookie("cookie")
	//判断是否有cookie,确定权限
	if err!=nil{
		c.JSON(200,gin.H{
			"message": "未登录，请登录账号",
		})
		c.Abort()
	} else {
		//将获取的cookie写入上下文
		c.Set("cookie",value)
		c.Next()
		//返回值，并打印
		v,_:=c.Get("next")
		fmt.Println(v)
	}
}




//定义主函数
func main(){
	r:=gin.Default()
	//登录界面
	r.POST("/login",login)
	//访问界面
	r.GET("/hello",cookie, func(c *gin.Context) {
		//打印欢迎用户
		cookie,_:=c.Get("cookie")
		str:=cookie.(string)
		c.String(200,"快去女装"+str)
		//设置next关键词，测试next函数
		c.Set("next","test next")
	})
	r.Run()
}