package usecases

import (
	"errors"
	"log"
	"skrik/internal/auth"
	"skrik/internal/entities"
	repository "skrik/internal/repository"

	"github.com/golang-jwt/jwt"
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
func (au *AuthUsecase) Register(user *entities.User) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Println("couldn't hash password. err: ", err)
		return "", err
	}
	user.Password = string(hashedPassword)
	/*user.Refresh_token, err = auth.GenerateRefreshToken(user.ID) //here's the refresh token generation functionality, but i can't properly test it so...
	if err != nil {
		return "", err
	}*/
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

func (au *AuthUsecase) CompareRefreshTokens(requestToken string) (string, error) {
	token, err := auth.ParseToken(requestToken)
	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}
	claims := token.Claims.(jwt.MapClaims)
	userid := uint(claims["userid"].(float64))
	user, err := au.repo.FindUserById(userid)
	if err != nil {
		return "", err
	}
	if user.Refresh_token != requestToken {
		return "", errors.New("refresh tokens doesn't match")
	}

	return auth.GenerateAccessToken(user.ID)
}
