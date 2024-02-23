package controller

import (
	"app/entity"
	"app/usecase"
)

type UserController struct {
	UserUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *UserController {
	return &UserController{
		UserUsecase: userUsecase,
	}
}

func (c *UserController) FindAll() ([]*entity.User, error) {
	return c.UserUsecase.FindAll()
}
