package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
func ConnectDB() (*gorm.DB,error){

	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %v", err)
	}
	
	dbUser := os.Getenv("DB_USERNAME")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")


	fmt.Printf("DB_USERNAME: %s\n", os.Getenv("DB_USERNAME"))
	fmt.Printf("DB_PASSWORD: %s\n", os.Getenv("DB_PASSWORD"))
	fmt.Printf("DB_NAME: %s\n", os.Getenv("DB_NAME"))
	fmt.Printf("DB_HOST: %s\n", os.Getenv("DB_HOST"))


	if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" {
        return nil, fmt.Errorf("database configuration environment variables are not set")
    }

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	
	if err != nil{
		return	nil,fmt.Errorf("không thể kết nối tới databases: %v",err)
	}
	return db,nil
}
