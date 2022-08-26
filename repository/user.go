package repository

type User struct {
	ID           string
	Name         string
	Email        string
	CreationTime string
}

type UserRepository interface {
	CreateUser(name, email string) (*User, error)
}
