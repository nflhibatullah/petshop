package city

import (
	"errors"
	"petshop/entity"

	"gorm.io/gorm"
)

type City interface {
	CreateCity(newCity entity.City) (entity.City, error)
	GetAllCity() ([]entity.City, error)
	GetCityByID(cityID int) (entity.City, error)
	UpdateCity(cityID int, updatedCity entity.City) (entity.City, error)
	DeletecCity(cityID int) (entity.City, error)
}

type cityRepository struct {
	db *gorm.DB
}

func NewCityRepository(db *gorm.DB) *cityRepository {
	return &cityRepository{db}
}

func (cr *cityRepository) CreateCity(newCity entity.City) (entity.City, error) {
	err := cr.db.Save(&newCity).Error

	if err != nil {
		return newCity, err
	}

	return newCity, nil
}

func (cr *cityRepository) GetAllCity() ([]entity.City, error) {
	cities := []entity.City{}

	err := cr.db.Find(&cities).Error

	if err != nil {
		return cities, err
	}

	return cities, err
}

func (cr *cityRepository) GetCityByID(cityID int) (entity.City, error) {
	city := entity.City{}

	err := cr.db.Where("id = ?", cityID).Find(&city).Error

	if err != nil || city.ID == 0 {
		return city, err
	}

	return city, err
}

func (cr *cityRepository) UpdateCity(cityID int, updatedCity entity.City) (entity.City, error) {
	city := entity.City{}

	err := cr.db.Where("id = ?", cityID).Find(&city).Error

	if err != nil || city.ID == 0 {
		return city, err
	}

	cr.db.Model(&city).Updates(&updatedCity)

	return updatedCity, err
}

func (cr *cityRepository) DeletecCity(cityID int) (entity.City, error) {
	city := entity.City{}

	err := cr.db.Where("id = ?", cityID).Find(&city).Error

	if err != nil || city.ID == 0 {
		return city, errors.New("")
	}

	cr.db.Delete(&city)

	return city, err
}
