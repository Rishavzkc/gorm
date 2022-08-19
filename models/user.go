package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	CompanyID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(id int, name string, email string) *User {
	return &User{
		CompanyID:    id,
		Name:  name,
		Email: email,
	}
}

func NewUserWithID(id int, name string, email string) *User {
	return &User{
		CompanyID:    id,
		Name:  name,
		Email: email,
	}
}

func (u *User) TableName() string {
	return "user"
}
