package usecase

import (
	"context"
	"errors"
	"shop/entity"
	"shop/repository"
	"shop/helper"
	"shop/logger"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface{
	Login(c context.Context,email, password string)(*entity.UserRegister, error)
	Register(c context.Context,userRegis entity.UserRegister) error
}

type authUseCase struct{
	AuthRepository  repository.AuthRepository
}

func NewAuthUseCase(authRepo repository.AuthRepository) AuthUseCase {
    return &authUseCase{AuthRepository : authRepo}
}

func (uc *authUseCase) Login(c context.Context,email,password string)(*entity.UserRegister, error){
	user, err := uc.AuthRepository.FindByEmail(c, email)
	if err != nil{
		return nil, err
	}
	if user == nil || bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password)) != nil {
		return nil, errors.New("invalide email or password")
	}
	return user, nil
}
func (uc *authUseCase) Register(c context.Context,userRegis entity.UserRegister) error{

	hashedPassword, err := helper.HashPassword(userRegis.Password)
	if err != nil {
		return err
	}
	existingUser, err := uc.AuthRepository.FindByEmail(c, userRegis.Email)
	if err != nil {
		return err // Lỗi khi tìm kiếm email
	}
	if existingUser != nil {
		 logger.Logger.Printf("email %s is already in use", userRegis.Email) // Trả lỗi nếu email đã tồn tại
		 return nil
	}

	user := entity.User{
        Name:     userRegis.Name,
        Email:    userRegis.Email,
        Password: hashedPassword,
        Phone:    &userRegis.Phone,
        IsAdmin:  userRegis.IsAdmin,
    }

	if err := uc.AuthRepository.CreateUser(c, user); err != nil {
		return err
	}

	return nil
}