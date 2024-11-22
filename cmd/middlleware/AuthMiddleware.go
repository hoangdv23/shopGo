package middlleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

var JwtSecret = os.Getenv("SECRET_KEY")

func AuthMiddleware() gin.HandlerFunc{
	return	func(c *gin.Context){
		// authheader := c.getHeader("Authorization")

	}

}