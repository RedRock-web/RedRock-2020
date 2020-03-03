package users

import (
	"RedRock-2020/0/database"
	"RedRock-2020/0/jwts"
	"RedRock-2020/0/response"
	"RedRock-2020/0/struct"
	"errors"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	f := BindJson(c)

	if IsRegiste(f.Username) {
		response.Error(c, 10003, "user exist!")
	} else {
		database.Insert(_struct.User{Username: f.Username, Password: f.Password}, "register insert record error!")
	}
	token := GetJwt(f, "register create jwt error!")
	response.OkWithData(c, gin.H{"token": token})
}

func IsRegiste(username string) bool {
	var user _struct.User
	database.G_db.Where("username = ?", username).First(&user)
	return user.ID != 0
}

func BindJson(c *gin.Context) (f _struct.LoginForm) {
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

func PasswdIsOk(f _struct.LoginForm) bool {
	var user _struct.User
	database.G_db.Where(_struct.User{
		Username: f.Username,
		Password: f.Password,
	}).First(&user)
	return user.ID != 0
}

func GetJwt(f _struct.LoginForm, errMsg string) string {
	j := jwts.NewJwt()
	token, err := j.Create(f, "redrock")
	if err != nil {
		errors.New(errMsg)
	}
	return token
}

func Modify(c *gin.Context) {
	m := _struct.ModifyForm{}
	if err := c.ShouldBindJSON(&m); err != nil {
		response.FormError(c)
		errors.New("bind json error when modify info!")
	}
	if AimsIsOk(m.Aims) {
		err := database.G_db.Model(&_struct.User{}).Where("username = ?", c.Keys["username"]).Update(m.Aims, m.Content).Error
		if err != nil {
			errors.New("modify info errors!")
		}
		response.Ok(c)
	} else {
		response.FormError(c)
	}
}

func AimsIsOk(aims string) bool {
	return aims == "gender" || aims == "nickname" || aims == "introduction"
}

func GetInfo(c *gin.Context) {
	username := c.Query("username")
	info := _struct.User{}
	err := database.G_db.Where("username = ?", username).Find(&info).Error
	if err != nil {
		errors.New("get info error!")
		return
	}
	response.OkWithData(c, gin.H{
		"gender":      info.Gender,
		"nickname":    info.Nickname,
		"uid":         info.Uid,
		"introdction": info.Introduction,
	})
}
