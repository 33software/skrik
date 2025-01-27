package entities

import (
	"gorm.io/gorm"
	"time"
)


type Room struct {
	gorm.Model
	Name string
}
type Message struct {
    gorm.Model
    Content   string    
    RoomID    uint      
    UserID    uint      
    SentAt    time.Time 
}

