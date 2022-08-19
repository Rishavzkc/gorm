package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	CompanyID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

func NewEmployee(id int, name string, department string) *Employee {
	return &Employee{
		CompanyID:         id,
		Name:       name,
		Department: department,
	}
}

func NewEmployeeWithID(id int, name string, department string) *Employee {
	return &Employee{
		CompanyID:         id,
		Name:       name,
		Department: department,
	}
}

func (e *Employee) TableName() string {
	return "employee"
}
