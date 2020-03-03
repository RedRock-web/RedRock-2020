package database

import (
	"RedRock-2020/aaa"
	"testing"
)

func TestInsert(t *testing.T) {
	G_db = Init()
	a := aaa.User{
		Username: "sfd",
		Password: "gwe",
	}

	Insert(a, "sdf")
}
