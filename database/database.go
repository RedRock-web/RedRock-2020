package database

import (
	"RedRock-2020/aaa"
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

	if G_db.HasTable(aaa.User{}) {
		G_db.AutoMigrate(aaa.User{})
	} else {
		if err = G_db.CreateTable(&aaa.User{}).Error; err != nil {
			//fmt.Println(err)
			errors.New("create table users error!")
		}
	}

	if G_db.HasTable(aaa.Infomation{}) {
		G_db.AutoMigrate(aaa.Infomation{})
	} else {
		if err = G_db.CreateTable(&aaa.Infomation{}).Error; err != nil {
			//fmt.Println(err)
			errors.New("create table Infomation error!")
		}
	}

	return G_db
}

func Insert(i interface{}, errMsg string) error {
	switch t := i.(type) {
	case aaa.User:
		u := i.(aaa.User)
		if err := G_db.Create(&u).Error; err != nil {
			errors.New(errMsg)
			return err
		}
	case aaa.Infomation:
		info := i.(aaa.Infomation)
		if err := G_db.Create(&info).Error; err != nil {
			errors.New(errMsg)
			return err
		}
	default:
		_ = t
	}
	return nil
}
