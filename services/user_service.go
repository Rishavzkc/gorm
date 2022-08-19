package services

import (
	"mysqlgin/interfaces"
	"mysqlgin/models"
)

type UserService struct {
	repository interfaces.UserRepository
}

func NewUserService(r interfaces.UserRepository) *UserService {
	return &UserService{
		repository: r,
	}
}

func (s *UserService) CreateUser(data models.User) (models.User, error) {
	use := models.NewUser(data.CompanyID, data.Name, data.Email)

	return s.repository.CreateUser(*use)
}

func (s *UserService) UpdateUser(data models.User) error {
	use := models.NewUserWithID(data.CompanyID, data.Name, data.Email)

	err := s.repository.UpdateUser(*use)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) DeleteUser(id int) (models.User, error) {
	return s.repository.DeleteUser(id)
}

func (s *UserService) ShowUser(id int) (models.User, error) {
	return s.repository.GetUserById(id)
}

func (s *UserService) ShowAllUsers() ([]models.User, error) {
	return s.repository.GetAllUser()
}
