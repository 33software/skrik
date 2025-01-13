package controllers

import (
	usecases "skrik/internal/usecases/user"
)

type UserController struct{
	usecase *usecases.UserUsecase
}

//creating new user controller object
func NewUserController (usecase *usecases.UserUsecase) *UserController {
	return &UserController{usecase: usecase}
}

