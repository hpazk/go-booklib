package user

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserServices interface {
	signUp(req *request) (User, error)
	FetchUsers() ([]User, error)
	FetchUserById(id int) (User, error)
	FetchUserByEmail(email string) (User, error)
	UpdateUser(id int, req *updateRequest) (User, error)
	DeleteUser(id int) error
}

type services struct {
	repository userRepository
}

func UserService(repository userRepository) *services {
	return &services{repository}
}

func (s *services) signUp(req *request) (User, error) {
	userReg := User{}
	userReg.Name = req.Name
	userReg.Address = req.Address
	// userReg.Photo = ""
	userReg.Email = req.Email
	// userReg.EmailVerifiedAt = time.Now()
	userReg.Password = req.Password
	// userReg.Role = ""
	userReg.CreatedAt = time.Now()
	userReg.UpdatedAt = time.Now()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return userReg, err
	}

	userReg.Password = string(hashedPassword)

	newUser, err := s.repository.Store(userReg)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *services) FetchUsers() ([]User, error) {
	var users []User
	users, _ = s.repository.Fetch()

	return users, nil
}

func (s *services) FetchUserById(id int) (User, error) {
	var user User
	user, _ = s.repository.GetById(id)

	return user, nil
}

func (s *services) FetchUserByEmail(email string) (User, error) {
	var user User
	user, _ = s.repository.GetByEmail(email)

	return user, nil
}

func (s *services) UpdateUser(id int, req *updateRequest) (User, error) {
	userReg := User{}
	userReg.ID = uint(id)
	userReg.Name = req.Name
	userReg.Address = req.Address
	// userReg.Photo = ""
	userReg.Email = req.Email
	// userReg.Role = ""
	userReg.CreatedAt = time.Now()
	userReg.UpdatedAt = time.Now()

	editUser, err := s.repository.Update(id, userReg)
	if err != nil {
		return editUser, err
	}

	return editUser, nil
}

func (s *services) DeleteUser(id int) error {
	if err := s.repository.Delete(id); err != nil {
		return errors.New("something went wrong")
	}

	return nil
}
