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
	f := BindJson(c)

	if IsRegiste(f.Username) {
		response.Error(c, 10003, "user exist!")
	} else {
		database.Insert(aaa.User{Username: f.Username, Password: f.Password}, "register insert record error!")
	}
	token := GetJwt(f, "register create jwt error!")
	response.OkWithData(c, gin.H{"token": token})
}

func IsRegiste(username string) bool {
	var user aaa.User
	database.G_db.Where("username = ?", username).First(&user)
	return user.ID != 0
}

func BindJson(c *gin.Context) (f aaa.LoginForm) {
	if err := c.ShouldBindJSON(&f); err != nil {
		response.FormError(c)
		//fmt.Println(err)
		errors.New("bind json error!")
		return
	}
	return f
}

func Login(c *gin.Context) {
	f := BindJson(c)
	if PasswdIsOk(f) {
		token := GetJwt(f, "get jwt error when login")
		response.OkWithData(c, token)
	} else {
		if IsRegiste(f.Username) {
			response.Error(c, 10004, "password error!")
		} else {
			response.Error(c, 10005, "unregistered!")
		}
	}
}

func PasswdIsOk(f aaa.LoginForm) bool {
	var user aaa.User
	database.G_db.Where(aaa.User{
		Username: f.Username,
		Password: f.Password,
	}).First(&user)
	return user.ID == 0
}

func GetJwt(f aaa.LoginForm, errMsg string) string {
	j := jwts.NewJwt()
	token, err := j.Create(f, "redrock")
	if err != nil {
		errors.New(errMsg)
	}
	return token
}
