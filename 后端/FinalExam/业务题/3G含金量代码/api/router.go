package api

import "github.com/gin-gonic/gin@v1.7.4"

//创建路由,返回路由引擎指针，供main调用，
//并包装不同的请求方法,api文件夹下其他的接口结合

func CreateRout()*gin.Engine{
	engine := gin.Default()
	return engine
}

func POST(relitivePath string,engine *gin.Engine,function func (context *gin.Context)){
	engine.POST(relitivePath, function)

}