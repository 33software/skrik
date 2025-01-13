package repository

import (
	"errors"
	"log"
	entities "skrik/internal/entities/user"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

//creating new user repository object
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser (user *entities.User) error{
	return ur.db.Create(user).Error
}
func (ur *UserRepository) FindUserById(id uint) (*entities.User, error){
	var user *entities.User
	if err := ur.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("user not found")
			return nil, err
		}
		log.Println("database error! err: ", err)
		return nil, err
	}
	return user, nil
}
/*func (ur *UserRepository) ChangePassword (user *entities.User) error{
	var foundUser *entities.User
	ur.db.First(&foundUser, user.ID)

	return 
}*/

func (ur *UserRepository) DeleteUser(id uint) error{
	var user *entities.User
	if id == 0 {
	log.Println("wrong id")
	return nil //add custom errors or smth
	}
	ur.db.Delete(&user, id)
	return nil
}