package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	CompanyID int        `json:"id"`
	Name      string     `json:"name"`
	Location  string     `json:"location"`
	Employees []Employee `gorm:"foreignkey:company_id"`
	//	Users     []User     `gorm:"many2many:companies_users;"`
}

func NewCompany(companyId int, name string, location string) *Company {
	return &Company{
		CompanyID: companyId,
		Name:      name,
		Location:  location,
	}
}

func NewCompanyWithID(companyId int, name string, location string) *Company {
	return &Company{
		CompanyID: companyId,
		Name:      name,
		Location:  location,
	}
}

func (c *Company) TableName() string {
	return "company"
}
