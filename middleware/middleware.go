package middleware

import (
	"RedRock-2020/jwts"
	"RedRock-2020/response"
	"github.com/gin-gonic/gin"
)

func AuthorityRequried() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if len(auth) < 7 {
			response.Error(c, 10005, "token error!")
			c.Abort()
		}
		token := auth[7 : len(auth)-1]
		j := jwts.NewJwt()
		f, err := j.Check(token, "redrock")
		if err != nil {
			response.Error(c, 10005, "token error!")
			c.Abort()
		}
		c.Set("username", f.Username)
		c.Next()
	}
}
