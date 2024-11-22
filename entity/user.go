package entity

import "time"

type User struct {
    ID              uint64     `gorm:"primaryKey;autoIncrement"`
    Name            string     `gorm:"type:varchar(255);not null"`
    Email           string     `gorm:"type:varchar(255);not null;unique"`
    Phone           *string    `gorm:"type:varchar(250)"`
    IsAdmin         int        `gorm:"column:isAdmin;default:0"`
    Password        string     `gorm:"type:varchar(255);not null"`
    Status          int        `gorm:"default:1"`
    CreatedAt       time.Time  `gorm:"autoCreateTime"`
    UpdatedAt       time.Time  `gorm:"autoUpdateTime"`
}


type UserRegister struct {
    ID       uint64 `json:"id"`
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password,omitempty" binding:"required"`
    Phone    string `json:"phone" binding:"required"` 
    IsAdmin  int    `json:"isAdmin"`
}

func (UserRegister) TableName() string {
    return "users"
}