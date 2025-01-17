package usecases

import (
	"errors"
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
		return "", entities.NewBadRequestError("empty or corrupted username/password")
	}
	user, err := au.repo.FindUserByUsername(username)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", entities.NewUnauthorizedError("incorrect password")
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
		return "", entities.NewInternalServerError("internal server error. debug: couldn't hash password")
	}
	user.Password = string(hashedPassword)
	/*user.Refresh_token, err = auth.GenerateRefreshToken(user.ID) //here's the refresh token generation functionality, but i can't properly test it so...
	if err != nil {
		return "", err
	}*/
	err = au.repo.CreateUser(user)
	if err != nil {
		return "", err
	}
	token, err := auth.GenerateAccessToken(user.ID)
	if err != nil {
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
		return "", entities.NewUnauthorizedError("refresh tokens doesn't match")
	}

	return auth.GenerateAccessToken(user.ID)
}
