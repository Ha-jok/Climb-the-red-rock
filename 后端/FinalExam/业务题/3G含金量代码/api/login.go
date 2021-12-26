package api

import (
	"FinalExam/业务题/3G含金量代码/service"
	"github.com/gin-gonic/gin@v1.7.4"
	"net/http"
)

//创建一个函数作为中间件,使用户发送用户名和密码,并判断用户名密码是否正确
func loginMiddleware (c *gin.Context){
	//用户发送信息
	username := c.PostForm("username")
	password := c.PostForm("password")
	//判断用户密码是否正确,并处理错误
	err,bool := service.Login(username,password)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"error" : err,
			"tip" : "请联系管理员",
		})
		return
	}
	if bool == true {
		c.JSON(http.StatusOK,gin.H{
			"msg" : "登录成功",
			"username" : username,
			"password" : password,
		})
	}else {
		c.String(http.StatusOK,"密码错误")
	}

}
//创建登录接口
func Login(engine *gin.Engine){
	POST("/login",engine,loginMiddleware)
}