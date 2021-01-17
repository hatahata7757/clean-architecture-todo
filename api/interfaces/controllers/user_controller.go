package controllers

import (
	"api/domain/model"
	"api/interfaces/database"
	"api/usecase/interactor"
)

// UserController is a type that includes the usecase.interactor.UserInteractor
type UserController struct {
	Interactor interactor.UserInteractor
}

// NewUserController method returns an UserController instance
func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: interactor.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// Create method is the controller for user registration
func (controller *UserController) Create(c Context) {
	u := model.User{}
	c.Bind(&u)
	user, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, user)
}
