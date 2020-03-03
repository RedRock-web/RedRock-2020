package users

import (
	"RedRock-2020/aaa"
	"RedRock-2020/database"
	"RedRock-2020/jwts"
	"RedRock-2020/response"
	"errors"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var f aaa.LoginForm
	if err := c.ShouldBindJSON(&f); err != nil {
		response.FormError(c)
		//fmt.Println(err)
		errors.New("bind json error!")
		return
	}
	if err := database.G_db.Create(&aaa.User{Username: f.Username, Password: f.Password}).Error; err != nil {
		errors.New("register insert record error!")
		return
	}

	j := jwts.NewJwt()
	token, err := j.Create(f, "redrock")
	if err != nil {
		errors.New("register create jwt error!")
		return
	}

	response.OkWithData(c, gin.H{"token": token})
}
