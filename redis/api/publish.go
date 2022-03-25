package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"redis/service"
)

func PuTheme(c *gin.Context){
	//获取参数，确定发布的主题及信息
	theme := c.Param("theme")
	message := c.PostForm("message")
	//调用service层的publish函数，读取信息
	service.Publish(theme,message)
	c.JSON(http.StatusOK,gin.H{
		"msg" : "ok",
	})
}
