package main

import (
	"gin-skill/app/dao"
	"gin-skill/bootstrap"
	"gin-skill/global"
)

func main() {
	bootstrap.InitConfig()
	bootstrap.InitLog()

	bootstrap.Migrate()
	dao.SetDefault(bootstrap.DB())
	defer bootstrap.CloseDB()

	bootstrap.InitValidator()
	global.App.Redis = bootstrap.InitRedis()

	bootstrap.RunServer()
}
