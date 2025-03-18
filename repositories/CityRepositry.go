package repositories

import (
	"x-app/database"
	"x-app/models"
)

func CreateCity(city *models.City) error {
	return database.DB.Create(city).Error
}

func GetCities() ([]models.City, error) {
	var cities []models.City
	err := database.DB.Find(&cities).Error
	return cities, err
}

func GetCity(city *models.City) error {
	err := database.DB.First(&city, city.ID).Error
	return err
}

func UpdateCity(city *models.City) error {
	return database.DB.Save(city).Error
}

func DeleteCity(city *models.City) error {
	return database.DB.Delete(city).Error
}
