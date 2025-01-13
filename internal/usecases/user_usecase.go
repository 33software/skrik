package usecases

import (
	"log"
	entities "skrik/internal/entities"
	repository "skrik/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo *repository.UserRepository
}

// creatig enw user usecase object
func NewUserUsecase(repo *repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uuc *UserUsecase) RegisterUser(newUser *entities.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		log.Println("failed to encrypt password! err: ", err)
	}
	newUser.Password = string(hashedPassword)

	return uuc.repo.CreateUser(newUser)
}
func (uuc *UserUsecase) DeleteUser(id uint) error {
	err := uuc.repo.DeleteUser(id)
	if err != nil {
		log.Println("error! err: ", err)
		return err
	} 
	return nil
}
