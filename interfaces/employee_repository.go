package interfaces

import "mysqlgin/models"

type EmployeeRepository interface {
	GetAllEmployee() ([]models.Employee, error)
	CreateEmployee(e models.Employee) (models.Employee, error)
	GetEmployeeById(id int) (models.Employee, error)
	DeleteEmployee(id int) (models.Employee, error)
	UpdateEmployee(e models.Employee) error
}
