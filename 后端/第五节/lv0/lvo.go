package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	//创建一个默认的路由引擎
	r:=gin.Default()
//利用string渲染
	r.GET("/hello",func(c *gin.Context){
		c.String(http.StatusOK,"hello")
	})
	r.Run(":8080")
}