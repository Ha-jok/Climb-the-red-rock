package service

import (
	"fmt"
	"redis/dao"
)

func Publish(theme,message string){
	//连接redis
	err,rdb := dao.Link()
	if err != nil {
		fmt.Println("LinkRedisError:",err)
	}
	rdb.Publish(theme,message)
	return
}
