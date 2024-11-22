package handler

import (
	"net/http"
	"shop/entity"
	"shop/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{
	AuthUseCase usecase.AuthUseCase
}

func NewAuthHandler(authUseCase usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		AuthUseCase: authUseCase,
	}
}


func (h *AuthHandler) LoginHandler(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "handle xử lý login",
	})
}
func (h *AuthHandler) Registerhandler(c *gin.Context){
	var UserRegister entity.UserRegister
	if err := c.ShouldBindJSON(&UserRegister); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Invalid input data",
			"details": err.Error(),
		})
	}

	err := h.AuthUseCase.Register(c, UserRegister)

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Registration failed",
            "details": err.Error(),
        })
        return
    }

	c.JSON(http.StatusOK, gin.H{
        "message": "Registration successful",
        "user": map[string]interface{}{
            "name":  UserRegister.Name,
            "email": UserRegister.Email,
			"phone" : UserRegister.Phone,
			"isAdmin" : UserRegister.IsAdmin,
        },
    })
}