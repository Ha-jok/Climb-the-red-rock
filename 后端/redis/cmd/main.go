package main

import (
	"github.com/gin-gonic/gin"
	"redis/api"
)

func main() {
	engine := gin.Default()

	//订阅接口，订阅主题为theme
	engine.GET("/subscribe/:theme",api.SubTheme)

	//发布接口，每个人都可以通过该接口在theme主题中发布内容。
	engine.POST("/publish/:theme",api.PuTheme)


	engine.Run()
}
