package main

import (
	"RedRock-2020/jwts"
)

func main() {
	j := jwts.NewJwt()
	//s, _ := j.Create(users.LoginForm{"Sdfgw", "Ge"}, "redrock")
	//fmt.Println(s)
	j.Check("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJyZWRyb2NrIiwiZXhwIjoiMTU4MzI3NzcxNiIsImlhdCI6IjE1ODMyNjY5MTYiLCJ1c2VybmFtZSI6IlNkZmd3IiwicGFzc3dvcmQiOiJHZSJ9.bNlhMyupiOwAnZ7z926CRXeW5dwjJSwKHpv3Pmm29Rk=", "redrock")
	//
	//if err != nil {
	//	fmt.Println(a)
	//} else {
	//	fmt.Println(err)
	//}
}
