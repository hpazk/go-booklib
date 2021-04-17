package user

type userRepository interface {
	Store(user User) (User, error)
	Fetch() ([]User, error)
	Update(id int, user User) (User, error)
	Delete(id int) error
	GetById(id int) (User, error)
	GetByEmail(email string) (User, error)
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
func (r *repository) GetById(id int) (User, error) {
	return usersStorage[id-1], nil
}

// Get User By Email
func (r *repository) GetByEmail(email string) (User, error) {
	userByEmail := User{}

	for _, v := range usersStorage {
		if v.Email == email {
			userByEmail = v
		}
	}
	return userByEmail, nil
}

// Update user
func (r *repository) Update(id int, user User) (User, error) {
	usersStorage[id-1] = user
	return usersStorage[id-1], nil
}

// Delete User
func (r *repository) Delete(id int) error {
	usersStorage[id-1] = usersStorage[len(usersStorage)-1]
	usersStorage[len(usersStorage)-1] = User{}
	usersStorage = usersStorage[:len(usersStorage)-1]
	return nil
}
