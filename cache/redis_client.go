package cache

import (
	"context"
	"fmt"
	"gg_web_tmpl/common/config"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
)

var (
	client *redis.Client
	Mocker redismock.ClientMock
)

func InitMockClient() {
	cli, mock := redismock.NewClientMock()
	client = cli
	Mocker = mock
}

func InitRedis(ctx context.Context) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.GetConf().Redis.RedisHost, config.GetConf().Redis.RedisPort),
		Password: config.GetConf().Redis.RedisPasswd, // no password set
		DB:       0,                                  // use default DB
	})
	_, err = client.Ping(ctx).Result()
	if err != nil {
		return
	}
	return err
}

func GetRedisClient() *redis.Client {
	return client
}
