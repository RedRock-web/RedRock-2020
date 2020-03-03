package main

import (
	"RedRock-2020/jwts"
	"fmt"
)

func main() {
	j := jwts.NewJwt()
	//a, _ := j.Create(users.LoginForm{"Sdfgw", "Ge"}, "redrock")
	//fmt.Println(a)
	s, err := j.Check("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJyZWRyb2NrIiwiZXhwIjoiMTU4MzIxMTk3NyIsImlhdCI6IjE1ODMyMDExNzciLCJ1c2VybmFtZSI6IlNkZmd3IiwicGFzc3dvcmQiOiJHZSJ9.JM50fM3wbKNrw9SWIil1k9gtFahWS7QF3xeQ15D4Z2k=", "redrock")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}
