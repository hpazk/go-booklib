package user

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserServices interface {
	signUp(req *request) (User, error)
	signIn(req *loginRequest) (User, error)
	UploadPhoto(user User, fileLocation string) (User, error)
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

	// TODO userExist

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

// TODO user-login
func (s *services) signIn(req *loginRequest) (User, error) {
	email := req.Email
	password := req.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("invalid password")
	}

	return user, nil
}

func (s *services) UploadPhoto(user User, fileLocation string) (User, error) {
	user.Photo = fileLocation

	editedUser, err := s.repository.Update(user)
	if err != nil {
		return user, errors.New("something went wrong")
	}

	return editedUser, nil
}

func (s *services) FetchUsers() ([]User, error) {
	var users []User
	users, _ = s.repository.Fetch()

	return users, nil
}

func (s *services) FetchUserById(id int) (User, error) {
	var user User
	user, _ = s.repository.FindById(id)

	return user, nil
}

func (s *services) FetchUserByEmail(email string) (User, error) {
	var user User
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

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

	editedUser, err := s.repository.Update(userReg)
	if err != nil {
		return editedUser, err
	}

	return editedUser, nil
}

func (s *services) DeleteUser(id int) error {
	if err := s.repository.Delete(id); err != nil {
		return errors.New("something went wrong")
	}

	return nil
}
