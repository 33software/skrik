package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string `gorm:"unique" json:"username"`
	Password      string `json:"password"`
	Refresh_token string `json:"refresh_token"`
}
