package main

import (
	"RedRock-2020/database"
	"RedRock-2020/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.G_db = database.Init()
	router.SetupRouter(r)
	r.Run()
}
