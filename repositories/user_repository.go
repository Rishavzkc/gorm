package repositories

import (
	"fmt"
	"mysqlgin/database"
	"mysqlgin/models"
)

type UserRepository struct {
}

func (s *UserRepository) GetAllUser() ([]models.User, error) {
	db := database.GetDatabase()
	var u []models.User
	err := db.Find(&u).Error
	if err != nil {
		return nil, fmt.Errorf("cannot find user: %v", err)
	}

	return u, err
}

func (s *UserRepository) CreateUser(u models.User) (models.User, error) {
	db := database.GetDatabase()
	err := db.Create(&u).Error

	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func (s *UserRepository) GetUserById(id int) (models.User, error) {
	db := database.GetDatabase()
	var u models.User
	err := db.First(&u, id).Error

	if err != nil {
		return models.User{}, fmt.Errorf("cannot find user by id: %v", err)
	}
	return u, nil
}

func (s *UserRepository) DeleteUser(id int) (models.User, error) {
	db := database.GetDatabase()
	n, err := s.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}

	err = db.Delete(&n).Error
	if err != nil {
		return models.User{}, fmt.Errorf("cannot delete : %v", err)
	}

	return n, nil
}

func (s *UserRepository) UpdateUser(u models.User) error {
	db := database.GetDatabase()

	err := db.Save(&u).Error
	if err != nil {
		return fmt.Errorf("cannot update : %v", err)
	}

	return nil
}
