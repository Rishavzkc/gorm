package interfaces

import "mysqlgin/models"

type UserRepository interface {
	GetAllUser() ([]models.User, error)
	CreateUser(u models.User) (models.User, error)
	GetUserById(id int) (models.User, error)
	DeleteUser(id int) (models.User, error)
	UpdateUser(u models.User) error
}


