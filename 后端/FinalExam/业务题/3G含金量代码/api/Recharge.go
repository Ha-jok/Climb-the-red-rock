package api

import (
	"FinalExam/业务题/3G含金量代码/service"
	"github.com/gin-gonic/gin@v1.7.4"
	"net/http"
)

//创建中间件函数，调用service的接口，接收用户发送的信息，返回响应报文
func RechargeMiddleware(c *gin.Context){
	//接收用户发送的信息
	username := c.PostForm("username")
	money := c.PostForm("money")
	//调用service中的change接口，接收错误和返回值
	err,b := service.Recharge(username,money)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"error" : err,
			"tip" : "请联系管理员",
		})
		return
	}
	if b == true {
		c.JSON(http.StatusOK,gin.H{
			"msg" : "修改成功",
			"newpassword" : "已充值"money,
		})
	}else {
		c.String(http.StatusOK,"充值失败")
	}
}
