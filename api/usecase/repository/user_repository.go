package repository

import "api/domain/model"

// UserRepository is an interface that defines the data manipulation behavior of the User structure
type UserRepository interface {
	Store(model.User) (model.User, error)
	FindById(id int) (model.User, error)
}
