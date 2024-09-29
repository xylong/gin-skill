package main

import (
	"gin-skill/app/dao"
	"gin-skill/bootstrap"
)

func main() {
	bootstrap.InitConfig()
	bootstrap.InitLog()

	bootstrap.Migrate()
	dao.SetDefault(bootstrap.DB())
	defer bootstrap.CloseDB()

	bootstrap.InitValidator()

	bootstrap.RunServer()
}
