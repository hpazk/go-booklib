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

func (r *repository) Store(user User) (User, error) {
	user.ID = uint(len(users)) + 1
	users = append(users, user)
	return user, nil
}

func (r *repository) Fetch() ([]User, error) {
	return users, nil
}

// func (r *repository) Update(id int) (User, error) {

// }
