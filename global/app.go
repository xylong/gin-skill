package global

import (
	"gin-skill/config"
	"github.com/redis/go-redis/v9"
)

type Application struct {
	//ConfigViper *viper.Viper

	Config config.Configuration

	Redis *redis.Client
}

var App = new(Application)
