package database

import (
	"RedRock-2020/struct"
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	G_db = Init()
	err := G_db.Model(&_struct.User{}).Where("username = ?", "sdf").Update("nickname", "SDf").Error
	fmt.Println(err)
}
