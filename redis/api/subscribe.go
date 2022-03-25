package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"redis/service"
)

func SubTheme (c *gin.Context){
	//获取参数，确定订阅的主题
	theme := c.Param("theme")
	//调用service层的subscribe函数，读取信息
	channel,payload := service.Subscribe(theme)
	c.JSON(http.StatusOK,gin.H{
		"theme" : channel,
		"message" : payload,
	})

}
