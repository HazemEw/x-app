package repositories

import (
	"x-app/database"
	"x-app/models"
)

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	return users, err
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := database.DB.Preload("City").Where("id = ?", id).First(&user).Error
	return &user, err
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err // Return nil if not found
	}
	return &user, nil
}

func DeleteUser(id uint) error {
	return database.DB.Delete(&models.User{}, id).Error
}
