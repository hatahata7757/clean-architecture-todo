package database

import "api/domain/model"

// UserRepository is a type that includes the SqlHandler interface
type UserRepository struct {
	SqlHandler
}

// FindById method returns a single record assciated with the id passed as an argument
func (repo *UserRepository) FindById(id int) (user model.User, err error) {
	if err = repo.Find(&user, id).Error; err != nil {
		return
	}
	return
}

// Store method inserts the User information passed as an argument into the DB
func (repo *UserRepository) Store(u model.User) (user model.User, err error) {
	if err = repo.Create(&u).Error; err != nil {
		return
	}
	user = u
	return
}
