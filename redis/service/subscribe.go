package service

import (
	"fmt"
	"redis/dao"
)

func Subscribe(theme string)(string,string){
	//连接redis
	err,redisDB := dao.Link()
	if err != nil {
		fmt.Println("RedisLinkError:",err)
	}
	//连接一个订阅通道
	pubsub := redisDB.Subscribe(theme)
	//获取订阅通道里面的信息
	_,err = pubsub.Receive()
	if err != nil {
		fmt.Println("pubsub.receive.error:",err)
	}
	//创建go语言通道，储存信息。
	ch := pubsub.Channel()

	//获取通道内信息
	for msg := range ch {
		return msg.Channel,msg.Payload
	}
	return "",""
}
