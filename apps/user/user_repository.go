package user

import "errors"

type userRepository interface {
	Store(user User) (User, error)
	Fetch() ([]User, error)
	Update(user User) (User, error)
	Delete(id int) error
	FindById(id int) (User, error)
	FindByEmail(email string) (User, error)
}

type UsersStorage []User

var usersStorage UsersStorage

type repository struct {
	// *gorm.DB
	usersStorage *UsersStorage
}

func UserRepository(db *UsersStorage) *repository {
	return &repository{db}
}

// Save New User
func (r *repository) Store(user User) (User, error) {
	user.ID = uint(len(usersStorage)) + 1
	usersStorage = append(usersStorage, user)
	return user, nil
}

// Get All Users
func (r *repository) Fetch() ([]User, error) {
	var users []User
	for _, v := range usersStorage {
		users = append(users, v)
	}

	return users, nil
}

// Get User by Id
func (r *repository) FindById(id int) (User, error) {
	return usersStorage[id-1], nil
}

// Get User By Email
func (r *repository) FindByEmail(email string) (User, error) {
	userByEmail := User{}

	for _, v := range usersStorage {
		if v.Email == email {
			userByEmail = v
		} else {
			return v, errors.New("user email not found")
		}
	}
	return userByEmail, nil
}

// Update user
func (r *repository) Update(user User) (User, error) {
	usersStorage[user.ID-1] = user
	return usersStorage[user.ID-1], nil
}

// Delete User
func (r *repository) Delete(id int) error {
	usersStorage[id-1] = usersStorage[len(usersStorage)-1]
	usersStorage[len(usersStorage)-1] = User{}
	usersStorage = usersStorage[:len(usersStorage)-1]
	return nil
}
