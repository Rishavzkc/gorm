package repositories

import (
	"fmt"
	"mysqlgin/database"
	"mysqlgin/models"
)

type CompanyRepository struct {
}

func (s *CompanyRepository) GetAllCompany() ([]models.Company, error) {
	db := database.GetDatabase()
	var c []models.Company
	err := db.Find(&c).Error
	if err != nil {
		return nil, fmt.Errorf("cannot find company: %v", err)
	}

	return c, err
}

func (s *CompanyRepository) CreateCompany(c models.Company, emp models.Employee) (models.Company, models.Employee, error) {
	db := database.GetDatabase()
	//var emp models.Employee

	// employee := models.Employee{
	// 	Name:       "abc",
	// 	CompanyID:  77,
	// 	Department: "IT",
	// }

	company := models.Company{
		Name:      c.Name,
		CompanyID: c.CompanyID,
		Location:  c.Location,
		Employees: []models.Employee{{
			Name:       emp.Name,
			CompanyID:  emp.CompanyID,
			Department: emp.Department,
		},
		},
	}

	err := db.Create(&company).Error

	if err != nil {
		return models.Company{}, emp, err
	}
	return c, emp, nil
}

func (s *CompanyRepository) GetCompanyById(id int) (models.Company, error) {
	db := database.GetDatabase()
	var c models.Company
	//err := db.First(&c, id).Error
	err := db.Preload("Employees").Where("company_id =? ", id).First(&c).Error
	if err != nil {
		return models.Company{}, fmt.Errorf("cannot find company by id: %v", err)
	}
	return c, nil
}

func (s *CompanyRepository) DeleteCompany(id int) (models.Company, error) {
	db := database.GetDatabase()
	n, err := s.GetCompanyById(id)
	if err != nil {
		return models.Company{}, err
	}

	err = db.Delete(&n).Error
	if err != nil {
		return models.Company{}, fmt.Errorf("cannot delete : %v", err)
	}

	return n, nil
}

func (s *CompanyRepository) UpdateCompany(c models.Company) error {
	db := database.GetDatabase()

	err := db.Save(&c).Error
	if err != nil {
		return fmt.Errorf("cannot update : %v", err)
	}

	return nil
}
