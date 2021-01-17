package interactor

import (
	"api/domain/model"
	"api/usecase/repository"
)

type UserInteractor struct {
	UserRepository repository.UserRepository
}

// Add は UserRepository を通して Store を実行するためのメソッドです。
func (interactor *UserInteractor) Add(u model.User) (user model.User, err error) {
	user, err = interactor.UserRepository.Store(u)
	if err != nil {
		return
	}
	return
}
