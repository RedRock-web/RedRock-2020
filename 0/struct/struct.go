package _struct

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username     string
	Password     string
	Uid          int
	Gender       int
	Nickname     string
	Introduction string
}

type LoginForm struct {
	Username string `json:"username" `
	Password string `json:"password" `
}

type ModifyForm struct {
	Aims    string `json:"aims"`
	Content string `json:"content"`
}
