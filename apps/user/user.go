package user

import "time"

type User struct {
	Name            string
	Address         string
	Photo           string
	Email           string
	EmailVerifiedAt time.Time
	Password        string
	Role            string
}
