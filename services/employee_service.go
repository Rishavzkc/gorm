package services

import (
	"mysqlgin/interfaces"
	"mysqlgin/models"
)

type EmployeeService struct {
	repository interfaces.EmployeeRepository
}

func NewEmployeeService(r interfaces.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		repository: r,
	}
}

func (s *EmployeeService) CreateEmployee(data models.Employee) (models.Employee, error) {
	empl := models.NewEmployee(data.CompanyID, data.Name, data.Department)

	return s.repository.CreateEmployee(*empl)
}

func (s *EmployeeService) UpdateEmployee(data models.Employee) error {
	empl := models.NewEmployeeWithID(data.CompanyID, data.Name, data.Department)

	err := s.repository.UpdateEmployee(*empl)

	if err != nil {
		return err
	}

	return nil
}

func (s *EmployeeService) DeleteEmployee(id int) (models.Employee, error) {
	return s.repository.DeleteEmployee(id)
}

func (s *EmployeeService) ShowEmployee(id int) (models.Employee, error) {
	return s.repository.GetEmployeeById(id)
}

func (s *EmployeeService) ShowAllEmployee() ([]models.Employee, error) {
	return s.repository.GetAllEmployee()
}
