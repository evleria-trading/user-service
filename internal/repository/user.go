package repository

type User interface {
}

type user struct {
}

func NewUserRepository() User {
	return &user{}
}
