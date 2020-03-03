package router

import (
	"RedRock-2020/middleware"
	"RedRock-2020/users"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/register", users.Register)
	r.POST("/login")
	r.GET("/find")
	r.POST("/modify", middleware.AuthorityRequried())
}
