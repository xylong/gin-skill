package main

import (
	"fmt"
	"gin-skill/app/dao"
	"gin-skill/bootstrap"
	"gin-skill/global"
	"gin-skill/routes"
	"gin-skill/utils"
)

func main() {
	bootstrap.InitConfig()

	utils.Migrate()
	dao.SetDefault(utils.DB())

	router := routes.InitRouter()

	if err := utils.InitTrans("zh"); err != nil {
		fmt.Println("初始化翻译器错误")
		return
	}

	router.Run(":" + global.App.Config.App.Port)
}
