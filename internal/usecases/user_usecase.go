package usecases

import (
	//entities "skrik/internal/entities"
	"skrik/internal/entities"
	repository "skrik/internal/repository"
)

type UserUsecase struct {
	repo *repository.UserRepository
}

// creatig enw user usecase object
func NewUserUsecase(repo *repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uuc *UserUsecase) DeleteUser(id uint) error {
	err := uuc.repo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
func (uuc *UserUsecase) GetUserByID(userid uint) (*entities.User, error) {
	if userid == 0 {
		return nil, entities.NewBadRequestError("empty or corrupted data. debug: empty userid")
	}
	return uuc.repo.FindUserById(userid)
}
