package main

import (
	"fmt"
	"gin-skill/dao"
	"gin-skill/routers"
	"gin-skill/utils"
)

func main() {
	utils.Migrate()
	dao.SetDefault(utils.DB())

	router := routers.InitRouter()

	if err := utils.InitTrans("zh"); err != nil {
		fmt.Println("初始化翻译器错误")
		return
	}

	router.Run(":8080")
}
