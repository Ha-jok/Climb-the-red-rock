package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	engine := gin.Default()
	engine.GET("/hello",func (c *gin.Context){
		c.String(http.StatusOK,"Hello World!")
	})
	fmt.Println("Hello World!")
	engine.Run(":8080")
}
