package repository

import (
	"context"
	"errors"
	"shop/entity"
	"fmt"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindByEmail(c context.Context, email string) (*entity.UserRegister,error)
	CreateUser(c context.Context,user entity.User) error
}

type AuthRepositoryImpl struct{
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepositoryImpl {
    return &AuthRepositoryImpl{db: db}
}

func (r *AuthRepositoryImpl) FindByEmail(c context.Context, email string) (*entity.UserRegister, error) {
	var user entity.UserRegister
	if err := r.db.WithContext(c).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Trả về nil nếu không tìm thấy người dùng
		}
		return nil, err // Trả về nil và lỗi nếu có lỗi khác xảy ra
	}
	return &user, nil // Trả về người dùng và nil cho error nếu tìm thấy
}

func (r *AuthRepositoryImpl) CreateUser(c context.Context,user entity.User) error {
	fmt.Printf("Inserting user: %+v\n", user)
	return r.db.Create(&user).Error
}