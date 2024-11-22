package cmd

import (
	"log"
	// "os"
	"shop/config"
	"shop/handler"
	"shop/repository"
	"shop/route"
	"shop/usecase"
	"shop/logger"
	"github.com/gin-gonic/gin"
)


func ActionStart() {

	logFile, err := logger.PrintLog()
	defer logFile.Close()
	
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	r := gin.Default()
	r.SetTrustedProxies(nil)
	// r.GET("/health", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status":  "success",
	// 		"message": "API is running",
	// 	})
	// })
	

	authRepo := repository.NewAuthRepository(db)
	authUseCase := usecase.NewAuthUseCase(authRepo)
	authHandler := handler.NewAuthHandler(authUseCase)
	route.Auth(r,authHandler)
	
  	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}