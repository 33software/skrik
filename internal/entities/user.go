package entities

import (
    "gorm.io/gorm"

)
type User struct {
    gorm.Model
    Username string
    Password string
    Refresh_token string

}
