package main

import (
	"gin-skill/dao"
	"gin-skill/routers"
	"gin-skill/utils"
)

func main() {
	utils.Migrate()
	dao.SetDefault(utils.DB())

	router := routers.InitRouter()

	router.Run(":8080")
}
