package main

import (
	"RedRock-2020/jwts"
	"RedRock-2020/users"
	"fmt"
)

func main() {
	j := jwts.NewJwt()
	s, _ := j.Create(users.LoginForm{"Sdfgw", "Ge"}, "egseg")
	fmt.Println(s)
}
