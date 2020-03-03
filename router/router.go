package router

import (
	"RedRock-2020/middleware"
	"RedRock-2020/users"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/register", users.Register)
	r.POST("/login", users.Login)
	r.GET("/get", users.GetInfo)
	r.POST("/modify", middleware.AuthorityRequried(), users.Modify)
}
