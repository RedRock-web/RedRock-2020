package database

import (
	"RedRock-2020/0/struct"
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var G_db *gorm.DB

func Init() *gorm.DB {
	G_db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/user?charset=utf8&parseTime=true")
	if err != nil {
		//fmt.Println(err)
		errors.New("database init error!")
	}

	if G_db.HasTable(_struct.User{}) {
		G_db.AutoMigrate(_struct.User{})
	} else {
		if err = G_db.CreateTable(&_struct.User{}).Error; err != nil {
			//fmt.Println(err)
			errors.New("create table users error!")
		}
	}

	return G_db
}

func Insert(i interface{}, errMsg string) error {
	u := i.(_struct.User)
	if err := G_db.Create(&u).Error; err != nil {
		errors.New(errMsg)
		return err
	}
	return nil
}
