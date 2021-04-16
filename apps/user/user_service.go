package user

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// TODO 1: create user services

func signUp(req *request) (User, error) {
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

	// TODO development: repository
	repository := new(userRepository)
	s := repo(userStorage{})
	fmt.Println("repos", repository)

	newUser, err := s.Store(userReg)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
