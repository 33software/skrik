package repository

import (
	"errors"
	entities "skrik/internal/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// creating new user repository object
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(user *entities.User) error {
	return ur.db.Create(user).Error
}
func (ur *UserRepository) FindUserById(id uint) (*entities.User, error) {
	var user *entities.User
	if err := ur.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entities.NewNotFoundError("user not found. debug: ")
		}
		return nil, err
	}
	return user, nil
}
func (ur *UserRepository) FindUserByUsername(username string) (*entities.User, error) {
	var user entities.User
	if username == "" {
		return nil, entities.NewBadRequestError("empty or corrupted data. debug: ")
	}
	err := ur.db.Where("Username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return nil, entities.NewNotFoundError("user not found. debug: ")
		}
		return nil, err
	}

	return &user, nil
}

/*func (ur *UserRepository) ChangePassword (user *entities.User) error{
	var foundUser *entities.User
	ur.db.First(&foundUser, user.ID)

	return
}*/

func (ur *UserRepository) DeleteUser(id uint) error {
	var user *entities.User
	if err := ur.db.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}
