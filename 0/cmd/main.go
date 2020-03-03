package main

import (
	"RedRock-2020/0/database"
	"RedRock-2020/0/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.G_db = database.Init()
	router.SetupRouter(r)
	r.Run()
}
