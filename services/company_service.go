package services

import (
	"mysqlgin/interfaces"
	"mysqlgin/models"
)

type CompanyService struct {
	repository interfaces.CompanyRepository
}

func NewCompanyService(r interfaces.CompanyRepository) *CompanyService {
	return &CompanyService{
		repository: r,
	}
}

func (s *CompanyService) CreateCompany(data models.Company, emp models.Employee) (models.Company, models.Employee, error) {
	comp := models.NewCompany(data.CompanyID, data.Name, data.Location)
	//employee := models.NewEmployee(emp.CompanyID, emp.Name, emp.Department)
	return s.repository.CreateCompany(*comp,emp)
}

func (s *CompanyService) UpdateCompany(data models.Company) error {
	comp := models.NewCompanyWithID(data.CompanyID, data.Name, data.Location)

	err := s.repository.UpdateCompany(*comp)

	if err != nil {
		return err
	}

	return nil
}

func (s *CompanyService) DeleteCompany(id int) (models.Company, error) {
	return s.repository.DeleteCompany(id)
}

func (s *CompanyService) ShowCompany(id int) (models.Company, error) {
	return s.repository.GetCompanyById(id)
}

func (s *CompanyService) ShowAllCompanies() ([]models.Company, error) {
	return s.repository.GetAllCompany()
}
