package services

import (
	"errors"
	"log"

	"github.com/gunjanmistry08/diary-app/internal/database"
	"github.com/gunjanmistry08/diary-app/internal/models"
)

func RegisterUser(username, password, email string) error {
	var user models.User
	result := database.DB.Where("username = ?", username).Or("email = ?", email).First(&user)
	if result.Error == nil {
		return errors.New("user already exists")
	}

	newUser := models.User{Username: username, Password: password, Email: email} // Hash in production
	if err := database.DB.Create(&newUser).Error; err != nil {
		log.Printf("failed to register user: %v", err)
		return err
	}
	return nil
}

func LoginUser(username, password string) error {
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		log.Printf("failed to login user: %v", err)
		return errors.New("invalid username or password")
	}

	if user.Password != password {
		return errors.New("invalid username or password")
	}
	return nil
}
