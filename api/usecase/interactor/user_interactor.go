package interactor

import (
	"api/domain/model"
	"api/usecase/repository"
)

// UserInteractor is a type includes the usecase.repository.UserRepository
type UserInteractor struct {
	UserRepository repository.UserRepository
}

// Add method is an interactor for user registration
func (interactor *UserInteractor) Add(u model.User) (user model.User, err error) {
	user, err = interactor.UserRepository.Store(u)
	if err != nil {
		return
	}
	return
}
