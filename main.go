package main

import (
	"gin-skill/routers"
	"gin-skill/utils"
)

func main() {
	utils.LoadDB()

	router := routers.InitRouter()

	router.Run(":8080")
}
