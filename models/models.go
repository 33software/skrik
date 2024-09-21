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
	gorm.Model `swaggerignore:"true"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type Response struct {
	gorm.Model `swaggerignore:"true"`
	Username   string `json:"username"`
	Email      string `json:"email"`
}

type ErrorResponse struct {
    Message string `json:"message"`
}
