package usecases

import (
	"errors"
	"log"
	"skrik/internal/auth"
	"skrik/internal/entities"
	repository "skrik/internal/repository"

	"golang.org/x/crypto/bcrypt"
)
type AuthUsecase struct {
	repo *repository.UserRepository
}

func NewAuthUsecase(repo *repository.UserRepository) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}

func (au *AuthUsecase) Authorize(username string, password string) (string, error) {
	if username == "" || password == "" {
		log.Println("corrupted username or password!")
		return "", errors.New("empty username or password")
	}
	user, err := au.repo.FindUserByUsername(username)
	if err != nil {
		log.Println("couldn't find user. err: ", err)
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			log.Println("wrong password!")
			return "", err
		}
		return "", err
	}
	token, err := auth.GenerateAccessToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
func (au *AuthUsecase) Register (user *entities.User) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Println("couldn't hash password. err: ", err)
		return "", err
	}
	user.Password = string(hashedPassword)
	err = au.repo.CreateUser(user)
	if err != nil {
		log.Println("couldn't create user. err: ", err)
		return "", err
	}
	token, err := auth.GenerateAccessToken(user.ID)
	if err != nil {
		log.Println("couldn't generate access token. err: ", err)
		return "", err
	}
	return token, nil
}