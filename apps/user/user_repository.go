package user

type userRepository interface {
	Fetch() ([]User, error)
	Update(id int) (User, error)
	Delete(id int)
	Store(user User) (User, error)
	GetByEmail(email string)
	GetById(id int)
}

type userStorage []User

var users userStorage

type repository struct {
	// *gorm.DB
	userStorage
}

func repo(db userStorage) *repository {
	return &repository{db}
}

// Save New User
func (r *repository) Store(user User) (User, error) {
	user.ID = uint(len(users)) + 1
	users = append(users, user)
	return user, nil
}

// Get All Users
func (r *repository) Fetch() ([]User, error) {
	return users, nil
}

// Get User
func (r *repository) GetById(id int) (User, error) {
	return users[id-1], nil
}

// Get User By Email
func (r *repository) GetByEmail(email string) (User, error) {
	userbyEmail := User{}

	for _, v := range users {
		if v.Email == email {
			userbyEmail = v
		}
	}
	return userbyEmail, nil
}

// Update user
func (r *repository) Update(id int) (User, error) {
	users[id-1].Name = "test"
	return users[id-1], nil
}

// Delete User
func (r *repository) Delete(id int) (User, error) {
	users[id-1] = users[len(users)-1]
	users[len(users)-1] = User{}
	users = users[:len(users)-1]
	return users[id-1], nil
}
