package global

import (
	"gin-skill/config"
)

type Application struct {
	//ConfigViper *viper.Viper

	Config config.Configuration
}

var App = new(Application)
