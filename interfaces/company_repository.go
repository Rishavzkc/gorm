package interfaces

import "mysqlgin/models"

type CompanyRepository interface {
	GetAllCompany() ([]models.Company, error)
	CreateCompany(c models.Company, emp models.Employee) (models.Company,models.Employee, error)
	GetCompanyById(id int) (models.Company, error)
	DeleteCompany(id int) (models.Company, error)
	UpdateCompany(c models.Company) error
}
