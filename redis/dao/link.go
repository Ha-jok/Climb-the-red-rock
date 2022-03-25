package dao

import "github.com/go-redis/redis"

func Link() (err error,rdb *redis.Client){
	//创建结构体，连接到redis
	redisDB := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//尝试连接，并处理错误。
	_, err = redisDB.Ping().Result()
	if err != nil {
		return err,redisDB
	}
	return nil,redisDB
}
