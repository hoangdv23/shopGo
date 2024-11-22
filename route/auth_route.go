package route

import (
	"shop/handler"

	"github.com/gin-gonic/gin"
)

func Auth(route *gin.Engine, authHandler *handler.AuthHandler){
	route.POST("/login", authHandler.LoginHandler)
	route.POST("/register", authHandler.Registerhandler)
}
