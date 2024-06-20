package redis

import (
	"context"
	"fsm/pkg/types"
	"github.com/go-redis/redis/v8"
)

func NewRedis(config *types.Config) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,  // Redis 服务器地址和端口
		Password: config.Redis.Password, // Redis 服务器密码，没有密码可以为空字符串
		DB:       0,                     // Redis 数据库编号，默认为0
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		panic("redis 连接错误:" + err.Error())
	}

	return client
}
