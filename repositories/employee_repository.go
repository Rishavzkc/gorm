package repositories

import (
	"fmt"
	"mysqlgin/database"
	"mysqlgin/models"
)

type EmployeeRepository struct {
}

func (s *EmployeeRepository) GetAllEmployee() ([]models.Employee, error) {
	db := database.GetDatabase()
	var e []models.Employee
	err := db.Find(&e).Error
	if err != nil {
		return nil, fmt.Errorf("cannot find employee: %v", err)
	}

	return e, err
}

func (s *EmployeeRepository) CreateEmployee(e models.Employee) (models.Employee, error) {
	db := database.GetDatabase()
	err := db.Create(&e).Error

	if err != nil {
		return models.Employee{}, err
	}
	return e, nil
}

func (s *EmployeeRepository) GetEmployeeById(id int) (models.Employee, error) {
	db := database.GetDatabase()
	var e models.Employee
	err := db.First(&e, id).Error

	if err != nil {
		return models.Employee{}, fmt.Errorf("cannot find employee by id: %v", err)
	}
	return e, nil
}

func (s *EmployeeRepository) DeleteEmployee(id int) (models.Employee, error) {
	db := database.GetDatabase()
	n, err := s.GetEmployeeById(id)
	if err != nil {
		return models.Employee{}, err
	}

	err = db.Delete(&n).Error
	if err != nil {
		return models.Employee{}, fmt.Errorf("cannot delete : %v", err)
	}

	return n, nil
}

func (s *EmployeeRepository) UpdateEmployee(e models.Employee) error {
	db := database.GetDatabase()

	err := db.Save(&e).Error
	if err != nil {
		return fmt.Errorf("cannot update : %v", err)
	}

	return nil
}
