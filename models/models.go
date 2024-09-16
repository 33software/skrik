package models

import (
	//"audio-stream-golang/config"
	"gorm.io/gorm"
)

/*type UserSchema struct {
	UserId   uint    `json:"userid"`
	Username string `json:"username"`
	Email    string `json:"email"`
}*/

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Username string
	Email    string `gorm:"unique"`
	Password string
}
