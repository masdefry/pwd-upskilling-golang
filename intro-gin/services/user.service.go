package services

import (
	"errors"
	"intro-gin/config"
	"intro-gin/models"
	"intro-gin/utils"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Register(username string, password string, role models.UserRole) (*models.User, error) {
	var existing models.User

	if err := config.DB.
		Where("username = ?", username).
		First(&existing).Error; err == nil {

		return nil, errors.New("username already exists")
	}
	
	hashedPassword, err := utils.HashPassword(password)

	if(err != nil){
		return nil, errors.New("Failed to hashed password")
	}

	user := &models.User{
		Username:  username,
		Password: hashedPassword,
		Role: role,
	}

	err = config.DB.Create(&user).Error

	if(err != nil){
		return nil, err;
	}

	return user, nil
}

func (s *UserService) Login(username string, password string) (*models.User, error){
	var user models.User

	// 1️⃣ Cari user berdasarkan username
	if err := config.DB.
		Where("username = ?", username).
		First(&user).Error; err != nil {

		// user tidak ditemukan
		return nil, errors.New("invalid username or password")
	}

	// 2️⃣ Compare password (plain vs hashed)
	if !utils.CheckPassword(user.Password, password) {
		return nil, errors.New("invalid username or password")
	}

	// 3️⃣ Login success
	return &user, nil
}