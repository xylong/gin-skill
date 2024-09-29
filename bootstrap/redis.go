package bootstrap

import (
	"context"
	"gin-skill/global"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

// InitRedis 初始化redis
func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.App.Config.Redis.Host + ":" + cast.ToString(global.App.Config.Redis.Port),
		Password: global.App.Config.Redis.Password, // no password set
		DB:       global.App.Config.Redis.DB,       // use default DB
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		zap.L().Error("Redis connect ping failed, err:", zap.Any("err", err))
		return nil
	}

	return client
}
