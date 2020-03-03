package aaa

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
}

type LoginForm struct {
	Username string `json:"username" `
	Password string `json:"password" `
}
